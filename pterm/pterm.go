/*!
 * Copyright 2023 Mario Finelli
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pterm

import (
	"github.com/pressly/goose/v3"
	"github.com/pterm/pterm"
)

// Returns a goose logger using our custom pterm-based logger.
func GoosePtermLogger() goose.Logger {
	return &goosePtermLogger{}
}

// This is the pterm type to implement all of the required functions for a
// custom goose logger.
type goosePtermLogger struct{}

func (*goosePtermLogger) Fatal(v ...interface{}) {
	pterm.Fatal.WithFatal(true).Println(v...)
}

func (*goosePtermLogger) Fatalf(format string, v ...interface{}) {
	pterm.Fatal.WithFatal(true).Printf(format, v...)
}

func (*goosePtermLogger) Print(v ...interface{}) {
	pterm.Info.Print(v...)
}

func (*goosePtermLogger) Println(v ...interface{}) {
	pterm.Info.Println(v...)
}

func (*goosePtermLogger) Printf(format string, v ...interface{}) {
	pterm.Info.Printf(format, v...)
}
