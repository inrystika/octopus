// MIT License
//
// Copyright (c) PCL. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE
//

package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	// init pprof server
	_ "net/http/pprof"

	"github.com/spf13/pflag"

	"k8s.io/apimachinery/pkg/util/wait"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog"

	"volcano.sh/volcano/cmd/scheduler/app"
	"volcano.sh/volcano/cmd/scheduler/app/options"
	// Import default actions/plugins.
	_ "volcano.sh/volcano/pkg/scheduler/actions"
	_ "volcano.sh/volcano/pkg/scheduler/plugins"

	// init assert
	_ "volcano.sh/volcano/pkg/scheduler/util/assert"
)

var logFlushFreq = pflag.Duration("log-flush-frequency", 5*time.Second, "Maximum number of seconds between log flushes")

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	klog.InitFlags(nil)

	s := options.NewServerOption()
	s.AddFlags(pflag.CommandLine)
	s.RegisterOptions()

	cliflag.InitFlags()
	if err := s.CheckOptionOrDie(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	go wait.Until(klog.Flush, *logFlushFreq, wait.NeverStop)
	defer klog.Flush()

	if err := app.Run(s); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
