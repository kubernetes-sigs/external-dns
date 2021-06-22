// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ipv4

import (
	"net"
	"syscall"
	"unsafe"

	"golang.org/x/net/internal/iana"
	"golang.org/x/net/internal/socket"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD

	"golang.org/x/sys/unix"
)

var (
	ctlOpts = [ctlMax]ctlOpt{
		ctlTTL:        {unix.IP_TTL, 1, marshalTTL, parseTTL},
		ctlPacketInfo: {unix.IP_PKTINFO, sizeofInetPktinfo, marshalPacketInfo, parsePacketInfo},
	}

	sockOpts = map[int]*sockOpt{
		ssoTOS:                {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_TOS, Len: 4}},
		ssoTTL:                {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_TTL, Len: 4}},
		ssoMulticastTTL:       {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_TTL, Len: 4}},
		ssoMulticastInterface: {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_IF, Len: unix.SizeofIPMreqn}, typ: ssoTypeIPMreqn},
		ssoMulticastLoopback:  {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTTL:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_RECVTTL, Len: 4}},
		ssoPacketInfo:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_PKTINFO, Len: 4}},
		ssoHeaderPrepend:      {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_HDRINCL, Len: 4}},
		ssoICMPFilter:         {Option: socket.Option{Level: iana.ProtocolReserved, Name: unix.ICMP_FILTER, Len: sizeofICMPFilter}},
		ssoJoinGroup:          {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup: {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======

>>>>>>> 5ce8c7613 (update vendored files)
	"golang.org/x/sys/unix"
)

var (
	ctlOpts = [ctlMax]ctlOpt{
		ctlTTL:        {unix.IP_TTL, 1, marshalTTL, parseTTL},
		ctlPacketInfo: {unix.IP_PKTINFO, sizeofInetPktinfo, marshalPacketInfo, parsePacketInfo},
	}

	sockOpts = map[int]*sockOpt{
<<<<<<< HEAD
		ssoTOS:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TOS, Len: 4}},
		ssoTTL:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TTL, Len: 4}},
		ssoMulticastTTL:       {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_TTL, Len: 4}},
		ssoMulticastInterface: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_IF, Len: sizeofIPMreqn}, typ: ssoTypeIPMreqn},
		ssoMulticastLoopback:  {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTTL:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_RECVTTL, Len: 4}},
		ssoPacketInfo:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_PKTINFO, Len: 4}},
		ssoHeaderPrepend:      {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_HDRINCL, Len: 4}},
		ssoICMPFilter:         {Option: socket.Option{Level: iana.ProtocolReserved, Name: sysICMP_FILTER, Len: sizeofICMPFilter}},
		ssoJoinGroup:          {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		ssoTOS:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TOS, Len: 4}},
		ssoTTL:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TTL, Len: 4}},
		ssoMulticastTTL:       {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_TTL, Len: 4}},
		ssoMulticastInterface: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_IF, Len: sizeofIPMreqn}, typ: ssoTypeIPMreqn},
		ssoMulticastLoopback:  {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTTL:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_RECVTTL, Len: 4}},
		ssoPacketInfo:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_PKTINFO, Len: 4}},
		ssoHeaderPrepend:      {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_HDRINCL, Len: 4}},
		ssoICMPFilter:         {Option: socket.Option{Level: iana.ProtocolReserved, Name: sysICMP_FILTER, Len: sizeofICMPFilter}},
		ssoJoinGroup:          {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
=======
		ssoTOS:                {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_TOS, Len: 4}},
		ssoTTL:                {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_TTL, Len: 4}},
		ssoMulticastTTL:       {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_TTL, Len: 4}},
		ssoMulticastInterface: {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_IF, Len: unix.SizeofIPMreqn}, typ: ssoTypeIPMreqn},
		ssoMulticastLoopback:  {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTTL:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_RECVTTL, Len: 4}},
		ssoPacketInfo:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_PKTINFO, Len: 4}},
		ssoHeaderPrepend:      {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_HDRINCL, Len: 4}},
		ssoICMPFilter:         {Option: socket.Option{Level: iana.ProtocolReserved, Name: unix.ICMP_FILTER, Len: sizeofICMPFilter}},
		ssoJoinGroup:          {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup: {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======

>>>>>>> 6b7ce455e (update vendored files)
	"golang.org/x/sys/unix"
)

var (
	ctlOpts = [ctlMax]ctlOpt{
		ctlTTL:        {unix.IP_TTL, 1, marshalTTL, parseTTL},
		ctlPacketInfo: {unix.IP_PKTINFO, sizeofInetPktinfo, marshalPacketInfo, parsePacketInfo},
	}

	sockOpts = map[int]*sockOpt{
<<<<<<< HEAD
		ssoTOS:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TOS, Len: 4}},
		ssoTTL:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TTL, Len: 4}},
		ssoMulticastTTL:       {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_TTL, Len: 4}},
		ssoMulticastInterface: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_IF, Len: sizeofIPMreqn}, typ: ssoTypeIPMreqn},
		ssoMulticastLoopback:  {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTTL:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_RECVTTL, Len: 4}},
		ssoPacketInfo:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_PKTINFO, Len: 4}},
		ssoHeaderPrepend:      {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_HDRINCL, Len: 4}},
		ssoICMPFilter:         {Option: socket.Option{Level: iana.ProtocolReserved, Name: sysICMP_FILTER, Len: sizeofICMPFilter}},
		ssoJoinGroup:          {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		ssoTOS:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TOS, Len: 4}},
		ssoTTL:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TTL, Len: 4}},
		ssoMulticastTTL:       {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_TTL, Len: 4}},
		ssoMulticastInterface: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_IF, Len: sizeofIPMreqn}, typ: ssoTypeIPMreqn},
		ssoMulticastLoopback:  {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTTL:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_RECVTTL, Len: 4}},
		ssoPacketInfo:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_PKTINFO, Len: 4}},
		ssoHeaderPrepend:      {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_HDRINCL, Len: 4}},
		ssoICMPFilter:         {Option: socket.Option{Level: iana.ProtocolReserved, Name: sysICMP_FILTER, Len: sizeofICMPFilter}},
		ssoJoinGroup:          {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
=======
		ssoTOS:                {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_TOS, Len: 4}},
		ssoTTL:                {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_TTL, Len: 4}},
		ssoMulticastTTL:       {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_TTL, Len: 4}},
		ssoMulticastInterface: {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_IF, Len: unix.SizeofIPMreqn}, typ: ssoTypeIPMreqn},
		ssoMulticastLoopback:  {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTTL:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_RECVTTL, Len: 4}},
		ssoPacketInfo:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_PKTINFO, Len: 4}},
		ssoHeaderPrepend:      {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_HDRINCL, Len: 4}},
		ssoICMPFilter:         {Option: socket.Option{Level: iana.ProtocolReserved, Name: unix.ICMP_FILTER, Len: sizeofICMPFilter}},
		ssoJoinGroup:          {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup: {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======

>>>>>>> 4d7e5ad26 (update vendored files)
	"golang.org/x/sys/unix"
)

var (
	ctlOpts = [ctlMax]ctlOpt{
		ctlTTL:        {unix.IP_TTL, 1, marshalTTL, parseTTL},
		ctlPacketInfo: {unix.IP_PKTINFO, sizeofInetPktinfo, marshalPacketInfo, parsePacketInfo},
	}

	sockOpts = map[int]*sockOpt{
<<<<<<< HEAD
		ssoTOS:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TOS, Len: 4}},
		ssoTTL:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TTL, Len: 4}},
		ssoMulticastTTL:       {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_TTL, Len: 4}},
		ssoMulticastInterface: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_IF, Len: sizeofIPMreqn}, typ: ssoTypeIPMreqn},
		ssoMulticastLoopback:  {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTTL:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_RECVTTL, Len: 4}},
		ssoPacketInfo:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_PKTINFO, Len: 4}},
		ssoHeaderPrepend:      {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_HDRINCL, Len: 4}},
		ssoICMPFilter:         {Option: socket.Option{Level: iana.ProtocolReserved, Name: sysICMP_FILTER, Len: sizeofICMPFilter}},
		ssoJoinGroup:          {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		ssoTOS:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TOS, Len: 4}},
		ssoTTL:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TTL, Len: 4}},
		ssoMulticastTTL:       {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_TTL, Len: 4}},
		ssoMulticastInterface: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_IF, Len: sizeofIPMreqn}, typ: ssoTypeIPMreqn},
		ssoMulticastLoopback:  {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTTL:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_RECVTTL, Len: 4}},
		ssoPacketInfo:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_PKTINFO, Len: 4}},
		ssoHeaderPrepend:      {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_HDRINCL, Len: 4}},
		ssoICMPFilter:         {Option: socket.Option{Level: iana.ProtocolReserved, Name: sysICMP_FILTER, Len: sizeofICMPFilter}},
		ssoJoinGroup:          {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
=======
		ssoTOS:                {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_TOS, Len: 4}},
		ssoTTL:                {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_TTL, Len: 4}},
		ssoMulticastTTL:       {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_TTL, Len: 4}},
		ssoMulticastInterface: {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_IF, Len: unix.SizeofIPMreqn}, typ: ssoTypeIPMreqn},
		ssoMulticastLoopback:  {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTTL:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_RECVTTL, Len: 4}},
		ssoPacketInfo:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_PKTINFO, Len: 4}},
		ssoHeaderPrepend:      {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.IP_HDRINCL, Len: 4}},
		ssoICMPFilter:         {Option: socket.Option{Level: iana.ProtocolReserved, Name: unix.ICMP_FILTER, Len: sizeofICMPFilter}},
		ssoJoinGroup:          {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:         {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup: {Option: socket.Option{Level: iana.ProtocolIP, Name: unix.MCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"golang.org/x/sys/unix"
)

var (
	ctlOpts = [ctlMax]ctlOpt{
		ctlTTL:        {sysIP_TTL, 1, marshalTTL, parseTTL},
		ctlPacketInfo: {sysIP_PKTINFO, sizeofInetPktinfo, marshalPacketInfo, parsePacketInfo},
	}

	sockOpts = map[int]*sockOpt{
		ssoTOS:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TOS, Len: 4}},
		ssoTTL:                {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_TTL, Len: 4}},
		ssoMulticastTTL:       {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_TTL, Len: 4}},
		ssoMulticastInterface: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_IF, Len: sizeofIPMreqn}, typ: ssoTypeIPMreqn},
		ssoMulticastLoopback:  {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTTL:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_RECVTTL, Len: 4}},
		ssoPacketInfo:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_PKTINFO, Len: 4}},
		ssoHeaderPrepend:      {Option: socket.Option{Level: iana.ProtocolIP, Name: sysIP_HDRINCL, Len: 4}},
		ssoICMPFilter:         {Option: socket.Option{Level: iana.ProtocolReserved, Name: sysICMP_FILTER, Len: sizeofICMPFilter}},
		ssoJoinGroup:          {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:         {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:   {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup: {Option: socket.Option{Level: iana.ProtocolIP, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
		ssoAttachFilter:       {Option: socket.Option{Level: unix.SOL_SOCKET, Name: unix.SO_ATTACH_FILTER, Len: unix.SizeofSockFprog}},
	}
)

func (pi *inetPktinfo) setIfindex(i int) {
	pi.Ifindex = int32(i)
}

func (gr *groupReq) setGroup(grp net.IP) {
	sa := (*sockaddrInet)(unsafe.Pointer(&gr.Group))
	sa.Family = syscall.AF_INET
	copy(sa.Addr[:], grp)
}

func (gsr *groupSourceReq) setSourceGroup(grp, src net.IP) {
	sa := (*sockaddrInet)(unsafe.Pointer(&gsr.Group))
	sa.Family = syscall.AF_INET
	copy(sa.Addr[:], grp)
	sa = (*sockaddrInet)(unsafe.Pointer(&gsr.Source))
	sa.Family = syscall.AF_INET
	copy(sa.Addr[:], src)
}
