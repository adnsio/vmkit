// Spin up Linux VMs with QEMU and Apple virtualization framework
// Copyright (C) 2021 VMKit Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package engine

import "errors"

var (
	ErrImageNotFound                = errors.New("engine: image not found")
	ErrInvalidChecksum              = errors.New("engine: invalid checksum")
	ErrInvalidDriver                = errors.New("engine: invalid driver")
	ErrInvalidSSHPort               = errors.New("engine: invalid ssh port")
	ErrUnsupportedArchitecture      = errors.New("engine: unsupported architecture")
	ErrVirtualMachineAlreadyExist   = errors.New("engine: virtual machine already exist")
	ErrVirtualMachineAlreadyRunning = errors.New("engine: virtual machine is already running")
	ErrVirtualMachineNotRunning     = errors.New("engine: virtual machine is not running")
)