// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package common

import (
	"context"

	"github.com/alecthomas/kingpin/v2"

	"github.com/gravitational/teleport/lib/auth"
	"github.com/gravitational/teleport/lib/service/servicecfg"
	"github.com/gravitational/teleport/tool/common/webauthnwin"
)

// webauthnwinCommand adapts webauthnwin.Command for tctl.
type webauthnwinCommand struct {
	impl *webauthnwin.Command
}

func (c *webauthnwinCommand) Initialize(app *kingpin.Application, _ *servicecfg.Config) {
	c.impl = webauthnwin.NewCommand(app)
}

func (c *webauthnwinCommand) TryRun(ctx context.Context, selectedCommand string, _ *auth.Client) (match bool, err error) {
	return c.impl.TryRun(ctx, selectedCommand)
}
