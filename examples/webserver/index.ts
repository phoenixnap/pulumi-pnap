// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import * as pnap from "@pulumi/pnap";

//let project = packet.getProject({name: "ci-project"}, { async: true });

let server = new pnap.Server("test-server", {
    hostname: "test-server",
    description: "This is test server for typescript",
    location: "PHX",
    os: "ubuntu/bionic",
    type: "s1.c1.tiny",
    sshKeys: ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDF9LdAFElNCi7JoWh6KUcchrJ2Gac1aqGRPpdZNowObpRtmiRCecAMb7bUgNAaNfcmwiQi7tos9TlnFgprIcfMWb8MSs3ABYHmBgqEEt3RWYf0fAc9CsIpJdMCUG28TPGTlRXCEUVNKgLMdcseAlJoGp1CgbHWIN65fB3he3kAZcfpPn5mapV0tsl2p+ZyuAGRYdn5dJv2RZDHUZBkOeUobwsij+weHCKAFmKQKtCP7ybgVHaQjAPrj8MGnk1jBbjDt5ws+Be+9JNjQJee9zCKbAOsIo3i+GcUIkrw5jxPU/RTGlWBcemPaKHdciSzGcjWboapzIy49qypQhZe1U75 user2@172.16.1.106"],
});

export let ip = server.hostname;
