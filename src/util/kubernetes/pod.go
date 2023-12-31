package kubernetes

import (
	"bufio"
	"context"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/utils/integer"
)

func GetLogs(instance *entity.Instance, log chan string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Log.Errorf("%s GetLog error", instance.Name)
		}
	}()
	config := &Config{
		ApiServer:  instance.Cluster.ApiServer,
		Token:      instance.Cluster.Token,
		KubeConfig: instance.Cluster.KubeConfig,
	}
	client, err := NewKubernetesClient(config)
	if err != nil {
		log <- "Create Kubernetes Client Error."
	}

	req := client.CoreV1().Pods(instance.Namespace).GetLogs(instance.Name, &v1.PodLogOptions{Follow: true})
	readCloser, err := req.Stream(context.TODO())
	if err != nil {
		log <- err.Error()
	}
	defer func(readCloser io.ReadCloser) {
		err := readCloser.Close()
		if err != nil {
			log <- err.Error()
		}
	}(readCloser)

	read := bufio.NewReader(readCloser)
	for {
		bytes, err := read.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log <- "complete"
			}
			log <- err.Error()
		}
		log <- string(bytes)
	}
}

func GetPods(instance *entity.Instance) ([]string, error) {
	config := &Config{
		ApiServer:  instance.Cluster.ApiServer,
		Token:      instance.Cluster.Token,
		KubeConfig: instance.Cluster.KubeConfig,
	}
	client, err := NewKubernetesClient(config)
	if err != nil {
		return nil, err
	}

	req, err := client.CoreV1().Pods(instance.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var podNames []string
	for _, item := range req.Items {
		podNames = append(podNames, item.Name)
	}
	return podNames, nil
}

func GetPodStatus(instance *entity.Instance) (*entity.PodStatus, error) {
	config := &Config{
		ApiServer:  instance.Cluster.ApiServer,
		Token:      instance.Cluster.Token,
		KubeConfig: instance.Cluster.KubeConfig,
	}
	client, err := NewKubernetesClient(config)
	if err != nil {
		return nil, err
	}
	pod, err := client.CoreV1().Pods(instance.Namespace).Get(context.TODO(), instance.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	podStatus := entity.PodStatus{}
	podStatus.Name = pod.Name
	podStatus.Phase = string(pod.Status.Phase)
	var podConditions []entity.PodCondition
	for _, condition := range pod.Status.Conditions {
		podCondition := entity.PodCondition{}
		podCondition.Status = string(condition.Status)
		podCondition.Message = condition.Message
		podCondition.Reason = condition.Reason
		podConditions = append(podConditions, podCondition)
	}
	podStatus.Conditions = podConditions
	return &podStatus, nil
}

func GetPodsStatus(instance *entity.Instance) ([]entity.PodStatus, error) {
	config := &Config{
		ApiServer:  instance.Cluster.ApiServer,
		Token:      instance.Cluster.Token,
		KubeConfig: instance.Cluster.KubeConfig,
	}
	client, err := NewKubernetesClient(config)
	if err != nil {
		return nil, err
	}
	podList, err := client.CoreV1().Pods(instance.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var podsStatus []entity.PodStatus
	for _, pod := range podList.Items {
		podStatus := entity.PodStatus{}
		podStatus.Name = pod.Name
		podStatus.Phase = string(pod.Status.Phase)
		var podConditions []entity.PodCondition
		for _, condition := range pod.Status.Conditions {
			podCondition := entity.PodCondition{}
			podCondition.Status = string(condition.Status)
			podCondition.Message = condition.Message
			podCondition.Reason = condition.Reason
			podConditions = append(podConditions, podCondition)
		}
		podStatus.Conditions = podConditions
		podsStatus = append(podsStatus, podStatus)
	}
	return podsStatus, nil
}

func GetPod(instance *entity.Instance) (*v1.Pod, error) {
	config := &Config{
		ApiServer:  instance.Cluster.ApiServer,
		Token:      instance.Cluster.Token,
		KubeConfig: instance.Cluster.KubeConfig,
	}
	client, err := NewKubernetesClient(config)
	if err != nil {
		return nil, err
	}
	return client.CoreV1().Pods(instance.Namespace).Get(context.TODO(), instance.Name, metav1.GetOptions{})
}

func GetStatus(instance *entity.Instance) (string, error) {
	config := &Config{
		ApiServer:  instance.Cluster.ApiServer,
		Token:      instance.Cluster.Token,
		KubeConfig: instance.Cluster.KubeConfig,
	}
	client, err := NewKubernetesClient(config)
	if err != nil {
		return "", err
	}
	pod, err := client.CoreV1().Pods(instance.Namespace).Get(context.TODO(), instance.Name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	for _, cond := range pod.Status.Conditions {
		if string(cond.Type) == constant.ContainersReady {
			if string(cond.Status) != constant.ConditionTrue {
				return "Unavailable", nil
			}
		} else if string(cond.Type) == constant.PodInitialized && string(cond.Status) != constant.ConditionTrue {
			return "Initializing", nil
		} else if string(cond.Type) == constant.PodReady {
			if string(cond.Status) != constant.ConditionTrue {
				return "Unavailable", nil
			}
			for _, containerState := range pod.Status.ContainerStatuses {
				if !containerState.Ready {
					return "Unavailable", nil
				}
			}
		} else if string(cond.Type) == constant.PodScheduled && string(cond.Status) != constant.ConditionTrue {
			return "Scheduling", nil
		}
	}
	return string(pod.Status.Phase), nil
}

func MaxContainerRestarts(instance *entity.Instance) (int, error) {
	config := &Config{
		ApiServer:  instance.Cluster.ApiServer,
		Token:      instance.Cluster.Token,
		KubeConfig: instance.Cluster.KubeConfig,
	}
	client, err := NewKubernetesClient(config)
	if err != nil {
		return -1, err
	}
	pod, err := client.CoreV1().Pods(instance.Namespace).Get(context.TODO(), instance.Name, metav1.GetOptions{})
	if err != nil {
		return -1, err
	}
	maxRestarts := 0
	for _, c := range pod.Status.ContainerStatuses {
		maxRestarts = integer.IntMax(maxRestarts, int(c.RestartCount))
	}
	return maxRestarts, nil
}

func GetPodInfo(instance *entity.Instance) (pods chan []*v1.Pod) {
	config := &Config{
		ApiServer:  instance.Cluster.ApiServer,
		Token:      instance.Cluster.Token,
		KubeConfig: instance.Cluster.KubeConfig,
	}
	client, err := NewKubernetesClient(config)
	if err != nil {
		return
	}
	factory := informers.NewSharedInformerFactory(client, 0)
	podInformer := factory.Core().V1().Pods()
	informer := podInformer.Informer()

	stopCh := make(chan struct{})
	defer close(stopCh)
	go factory.Start(stopCh)

	if !cache.WaitForCacheSync(stopCh, informer.HasSynced) {
		logger.Log.Errorf("Failed to sync pod info")
		return
	}

	//informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
	//	AddFunc: func(obj interface{}) {
	//		pod := obj.(*v1.Pod)
	//		pod.Status
	//	},
	//	UpdateFunc: func(oldObj, newObj interface{}) {
	//		pod := newObj.(*v1.Pod)
	//	},
	//	DeleteFunc: func(obj interface{}) {
	//		pod := obj.(*v1.Pod)
	//	},
	//})
	podLister := podInformer.Lister()
	podList, err := podLister.Pods(instance.Namespace).List(labels.Everything())
	if err != nil {
		return
	}
	pods <- podList
	<-stopCh
	return
}
