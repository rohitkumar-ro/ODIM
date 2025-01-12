#!/bin/bash
# (C) Copyright [2020] Hewlett Packard Enterprise Development LP
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may
# not use this file except in compliance with the License. You may obtain
# a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.

# Script is for generating certificate and private key
# for Client mode connection usage only


export configFolder="/etc/odimra_config"
export logFolder="/var/log/odimra_logs"

cd /bin

for i in api account_session aggregation event systems task fabrics telemetry managers update
do
        systemctl enable $i
        while [ $? -ne 0 ]
        do
                systemctl enable $i
        done
done

while true; do
	sleep 5s
done
