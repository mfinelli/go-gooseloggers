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

package zerolog

import (
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
)

// Returns a goose logger using our custom zerolog-based logger.
func GooseZerologLogger(l *zerolog.Logger) goose.Logger {
	return &gooseZerologLogger{l: l}
}

// This is the zerolog type to implement all of the required functions for a
// custom goose logger.
type gooseZerologLogger struct {
	l *zerolog.Logger
}

func (l *gooseZerologLogger) Fatal(v ...interface{}) {
	l.l.Fatal().Msg(fmt.Sprintln(v...))
}

func (l *gooseZerologLogger) Fatalf(format string, v ...interface{}) {
	l.l.Fatal().Msg(fmt.Sprintf(format, v...))
}

func (l *gooseZerologLogger) Print(v ...interface{}) {
	l.l.Info().Msg(fmt.Sprint(v...))
}

func (l *gooseZerologLogger) Println(v ...interface{}) {
	l.l.Info().Msg(fmt.Sprintln(v...))
}

func (l *gooseZerologLogger) Printf(format string, v ...interface{}) {
	l.l.Info().Msg(fmt.Sprintf(format, v...))
}
