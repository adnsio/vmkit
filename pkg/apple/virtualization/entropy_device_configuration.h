// Virtual Machine manager that supports QEMU and Apple virtualization framework on macOS
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

#ifndef entropy_device_configuration_h
#define entropy_device_configuration_h

#import <Virtualization/Virtualization.h>

void *VZVirtioEntropyDeviceConfiguration_init();

#endif /* entropy_device_configuration_h */
