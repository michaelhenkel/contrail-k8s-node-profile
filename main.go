/*
Copyright 2016 Juniper

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

*/
package main

import (
        "time"
	"fmt"
        "os"
        "io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)


func main(){
   err := createConfig()
   if err != nil {
     panic(err.Error())
   }
}

func createConfig() error{
        return retry(10, time.Second, func() error {
          hostname, err := os.Hostname()
	  if err != nil {
		panic(err.Error())
	  }
          nameSpaceByte, err := ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/namespace")
	  if err != nil {
		panic(err.Error())
	  }
          nameSpace := string(nameSpaceByte)
	  config, err := rest.InClusterConfig()
	  if err != nil {
		panic(err.Error())
	  }
	  clientset, err := kubernetes.NewForConfig(config)
	  if err != nil {
		panic(err.Error())
	  }
          currentPod, err := clientset.CoreV1().Pods(nameSpace).Get(
                              hostname, metav1.GetOptions{})
          if err != nil {
            return err
          }
          currentNodeName := currentPod.Spec.NodeName
          currentNode, err := clientset.CoreV1().Nodes().Get(
                               currentNodeName, metav1.GetOptions{})
          if err != nil {
            return err
          }
          var nodeProfile string
          for labelKey, labelValue := range currentNode.Labels{
              if labelKey == "opencontrail.org/profile"{
                  nodeProfile = labelValue
              }
          }
          if nodeProfile != "" {
              configMapClient := clientset.CoreV1().ConfigMaps(nameSpace)
              profileCm, err := configMapClient.Get(nodeProfile, metav1.GetOptions{})
              if err != nil {
                return err
              }
              if len(profileCm.Data) > 0 {
                  for key, value := range(profileCm.Data){
                      fmt.Printf("%s=%s\n", key, value)
                  }
              }
          }
          return nil
        })
}
