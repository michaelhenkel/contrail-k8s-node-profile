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
	"net"
        "os"
	"strings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
        "k8s.io/api/core/v1"
        "gopkg.in/yaml.v2"
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
          fmt.Println("hostname:", hostname)
         
	  config, err := rest.InClusterConfig()
	  if err != nil {
		panic(err.Error())
	  }
	  clientset, err := kubernetes.NewForConfig(config)
	  if err != nil {
		panic(err.Error())
	  }
          nodeList, err := clientset.CoreV1().Nodes().List(
                    metav1.ListOptions{LabelSelector: "node-role.kubernetes.io/master=",})
          if err != nil {
            return err
          }

          configMapClient := clientset.CoreV1().ConfigMaps("contrail")
          cm, err := configMapClient.Get("contrailcontrollernodes", metav1.GetOptions{})
          if err != nil {
            configMapClient.Create(configMap)
            fmt.Printf("Created %s\n", cm.Name)
          } else {
            configMapClient.Update(configMap)
            fmt.Printf("Updated %s\n", cm.Name)
          }
          return nil
        })
}
