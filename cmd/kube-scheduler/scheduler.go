/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/spf13/pflag"

	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	_ "k8s.io/kubernetes/pkg/util/prometheusclientgo" // load all the prometheus client-go plugins
	_ "k8s.io/kubernetes/pkg/version/prometheus"      // for version metric registration
)

func main() {
	rand.Seed(time.Now().UnixNano())

	/**
	 ***********************
	  使用了cobra第三方框架，创建cobra.Command对象。
	  创建时最核心的是重写 Run 方法。

	  Run方法会在command执行Execute方法时被调用

	  command对象的execute方法使用了模板方法模式，
	  内部只搭建了执行的架子，比如：run之前，run，run之后。
	  具体实现留给子类重写
	 ***********************
	 */
	command := app.NewSchedulerCommand()

	// TODO: once we switch everything over to Cobra commands, we can go back to calling
	// utilflag.InitFlags() (by removing its pflag.Parse() call). For now, we have to set the
	// normalize func and add the go flag set by hand.
	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	// utilflag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	/**
	***********************
	开始执行Execute方法，
	内部会调用Run方法
	***********************
	*/
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
