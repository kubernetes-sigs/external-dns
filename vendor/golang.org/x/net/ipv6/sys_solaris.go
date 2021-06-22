// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ipv6

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
		ctlTrafficClass: {unix.IPV6_TCLASS, 4, marshalTrafficClass, parseTrafficClass},
		ctlHopLimit:     {unix.IPV6_HOPLIMIT, 4, marshalHopLimit, parseHopLimit},
		ctlPacketInfo:   {unix.IPV6_PKTINFO, sizeofInet6Pktinfo, marshalPacketInfo, parsePacketInfo},
		ctlNextHop:      {unix.IPV6_NEXTHOP, sizeofSockaddrInet6, marshalNextHop, parseNextHop},
		ctlPathMTU:      {unix.IPV6_PATHMTU, sizeofIPv6Mtuinfo, marshalPathMTU, parsePathMTU},
	}

	sockOpts = map[int]*sockOpt{
		ssoTrafficClass:        {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_TCLASS, Len: 4}},
		ssoHopLimit:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_UNICAST_HOPS, Len: 4}},
		ssoMulticastInterface:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_IF, Len: 4}},
		ssoMulticastHopLimit:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_HOPS, Len: 4}},
		ssoMulticastLoopback:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTrafficClass: {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVTCLASS, Len: 4}},
		ssoReceiveHopLimit:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVHOPLIMIT, Len: 4}},
		ssoReceivePacketInfo:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVPKTINFO, Len: 4}},
		ssoReceivePathMTU:      {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVPATHMTU, Len: 4}},
		ssoPathMTU:             {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_PATHMTU, Len: sizeofIPv6Mtuinfo}},
		ssoChecksum:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_CHECKSUM, Len: 4}},
		ssoICMPFilter:          {Option: socket.Option{Level: iana.ProtocolIPv6ICMP, Name: unix.ICMP6_FILTER, Len: sizeofICMPv6Filter}},
		ssoJoinGroup:           {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:          {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======

	"golang.org/x/sys/unix"
>>>>>>> 5ce8c7613 (update vendored files)
)

var (
	ctlOpts = [ctlMax]ctlOpt{
		ctlTrafficClass: {unix.IPV6_TCLASS, 4, marshalTrafficClass, parseTrafficClass},
		ctlHopLimit:     {unix.IPV6_HOPLIMIT, 4, marshalHopLimit, parseHopLimit},
		ctlPacketInfo:   {unix.IPV6_PKTINFO, sizeofInet6Pktinfo, marshalPacketInfo, parsePacketInfo},
		ctlNextHop:      {unix.IPV6_NEXTHOP, sizeofSockaddrInet6, marshalNextHop, parseNextHop},
		ctlPathMTU:      {unix.IPV6_PATHMTU, sizeofIPv6Mtuinfo, marshalPathMTU, parsePathMTU},
	}

	sockOpts = map[int]*sockOpt{
<<<<<<< HEAD
		ssoTrafficClass:        {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_TCLASS, Len: 4}},
		ssoHopLimit:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_UNICAST_HOPS, Len: 4}},
		ssoMulticastInterface:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_IF, Len: 4}},
		ssoMulticastHopLimit:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_HOPS, Len: 4}},
		ssoMulticastLoopback:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTrafficClass: {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVTCLASS, Len: 4}},
		ssoReceiveHopLimit:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVHOPLIMIT, Len: 4}},
		ssoReceivePacketInfo:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPKTINFO, Len: 4}},
		ssoReceivePathMTU:      {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPATHMTU, Len: 4}},
		ssoPathMTU:             {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_PATHMTU, Len: sizeofIPv6Mtuinfo}},
		ssoChecksum:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_CHECKSUM, Len: 4}},
		ssoICMPFilter:          {Option: socket.Option{Level: iana.ProtocolIPv6ICMP, Name: sysICMP6_FILTER, Len: sizeofICMPv6Filter}},
		ssoJoinGroup:           {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:          {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		ssoTrafficClass:        {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_TCLASS, Len: 4}},
		ssoHopLimit:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_UNICAST_HOPS, Len: 4}},
		ssoMulticastInterface:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_IF, Len: 4}},
		ssoMulticastHopLimit:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_HOPS, Len: 4}},
		ssoMulticastLoopback:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTrafficClass: {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVTCLASS, Len: 4}},
		ssoReceiveHopLimit:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVHOPLIMIT, Len: 4}},
		ssoReceivePacketInfo:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPKTINFO, Len: 4}},
		ssoReceivePathMTU:      {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPATHMTU, Len: 4}},
		ssoPathMTU:             {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_PATHMTU, Len: sizeofIPv6Mtuinfo}},
		ssoChecksum:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_CHECKSUM, Len: 4}},
		ssoICMPFilter:          {Option: socket.Option{Level: iana.ProtocolIPv6ICMP, Name: sysICMP6_FILTER, Len: sizeofICMPv6Filter}},
		ssoJoinGroup:           {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:          {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
=======
		ssoTrafficClass:        {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_TCLASS, Len: 4}},
		ssoHopLimit:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_UNICAST_HOPS, Len: 4}},
		ssoMulticastInterface:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_IF, Len: 4}},
		ssoMulticastHopLimit:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_HOPS, Len: 4}},
		ssoMulticastLoopback:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTrafficClass: {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVTCLASS, Len: 4}},
		ssoReceiveHopLimit:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVHOPLIMIT, Len: 4}},
		ssoReceivePacketInfo:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVPKTINFO, Len: 4}},
		ssoReceivePathMTU:      {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVPATHMTU, Len: 4}},
		ssoPathMTU:             {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_PATHMTU, Len: sizeofIPv6Mtuinfo}},
		ssoChecksum:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_CHECKSUM, Len: 4}},
		ssoICMPFilter:          {Option: socket.Option{Level: iana.ProtocolIPv6ICMP, Name: unix.ICMP6_FILTER, Len: sizeofICMPv6Filter}},
		ssoJoinGroup:           {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:          {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======

	"golang.org/x/sys/unix"
>>>>>>> 6b7ce455e (update vendored files)
)

var (
	ctlOpts = [ctlMax]ctlOpt{
		ctlTrafficClass: {unix.IPV6_TCLASS, 4, marshalTrafficClass, parseTrafficClass},
		ctlHopLimit:     {unix.IPV6_HOPLIMIT, 4, marshalHopLimit, parseHopLimit},
		ctlPacketInfo:   {unix.IPV6_PKTINFO, sizeofInet6Pktinfo, marshalPacketInfo, parsePacketInfo},
		ctlNextHop:      {unix.IPV6_NEXTHOP, sizeofSockaddrInet6, marshalNextHop, parseNextHop},
		ctlPathMTU:      {unix.IPV6_PATHMTU, sizeofIPv6Mtuinfo, marshalPathMTU, parsePathMTU},
	}

	sockOpts = map[int]*sockOpt{
<<<<<<< HEAD
		ssoTrafficClass:        {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_TCLASS, Len: 4}},
		ssoHopLimit:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_UNICAST_HOPS, Len: 4}},
		ssoMulticastInterface:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_IF, Len: 4}},
		ssoMulticastHopLimit:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_HOPS, Len: 4}},
		ssoMulticastLoopback:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTrafficClass: {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVTCLASS, Len: 4}},
		ssoReceiveHopLimit:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVHOPLIMIT, Len: 4}},
		ssoReceivePacketInfo:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPKTINFO, Len: 4}},
		ssoReceivePathMTU:      {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPATHMTU, Len: 4}},
		ssoPathMTU:             {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_PATHMTU, Len: sizeofIPv6Mtuinfo}},
		ssoChecksum:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_CHECKSUM, Len: 4}},
		ssoICMPFilter:          {Option: socket.Option{Level: iana.ProtocolIPv6ICMP, Name: sysICMP6_FILTER, Len: sizeofICMPv6Filter}},
		ssoJoinGroup:           {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:          {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		ssoTrafficClass:        {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_TCLASS, Len: 4}},
		ssoHopLimit:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_UNICAST_HOPS, Len: 4}},
		ssoMulticastInterface:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_IF, Len: 4}},
		ssoMulticastHopLimit:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_HOPS, Len: 4}},
		ssoMulticastLoopback:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTrafficClass: {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVTCLASS, Len: 4}},
		ssoReceiveHopLimit:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVHOPLIMIT, Len: 4}},
		ssoReceivePacketInfo:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPKTINFO, Len: 4}},
		ssoReceivePathMTU:      {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPATHMTU, Len: 4}},
		ssoPathMTU:             {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_PATHMTU, Len: sizeofIPv6Mtuinfo}},
		ssoChecksum:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_CHECKSUM, Len: 4}},
		ssoICMPFilter:          {Option: socket.Option{Level: iana.ProtocolIPv6ICMP, Name: sysICMP6_FILTER, Len: sizeofICMPv6Filter}},
		ssoJoinGroup:           {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:          {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
=======
		ssoTrafficClass:        {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_TCLASS, Len: 4}},
		ssoHopLimit:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_UNICAST_HOPS, Len: 4}},
		ssoMulticastInterface:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_IF, Len: 4}},
		ssoMulticastHopLimit:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_HOPS, Len: 4}},
		ssoMulticastLoopback:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTrafficClass: {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVTCLASS, Len: 4}},
		ssoReceiveHopLimit:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVHOPLIMIT, Len: 4}},
		ssoReceivePacketInfo:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVPKTINFO, Len: 4}},
		ssoReceivePathMTU:      {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVPATHMTU, Len: 4}},
		ssoPathMTU:             {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_PATHMTU, Len: sizeofIPv6Mtuinfo}},
		ssoChecksum:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_CHECKSUM, Len: 4}},
		ssoICMPFilter:          {Option: socket.Option{Level: iana.ProtocolIPv6ICMP, Name: unix.ICMP6_FILTER, Len: sizeofICMPv6Filter}},
		ssoJoinGroup:           {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:          {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======

	"golang.org/x/sys/unix"
>>>>>>> 4d7e5ad26 (update vendored files)
)

var (
	ctlOpts = [ctlMax]ctlOpt{
		ctlTrafficClass: {unix.IPV6_TCLASS, 4, marshalTrafficClass, parseTrafficClass},
		ctlHopLimit:     {unix.IPV6_HOPLIMIT, 4, marshalHopLimit, parseHopLimit},
		ctlPacketInfo:   {unix.IPV6_PKTINFO, sizeofInet6Pktinfo, marshalPacketInfo, parsePacketInfo},
		ctlNextHop:      {unix.IPV6_NEXTHOP, sizeofSockaddrInet6, marshalNextHop, parseNextHop},
		ctlPathMTU:      {unix.IPV6_PATHMTU, sizeofIPv6Mtuinfo, marshalPathMTU, parsePathMTU},
	}

	sockOpts = map[int]*sockOpt{
<<<<<<< HEAD
		ssoTrafficClass:        {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_TCLASS, Len: 4}},
		ssoHopLimit:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_UNICAST_HOPS, Len: 4}},
		ssoMulticastInterface:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_IF, Len: 4}},
		ssoMulticastHopLimit:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_HOPS, Len: 4}},
		ssoMulticastLoopback:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTrafficClass: {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVTCLASS, Len: 4}},
		ssoReceiveHopLimit:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVHOPLIMIT, Len: 4}},
		ssoReceivePacketInfo:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPKTINFO, Len: 4}},
		ssoReceivePathMTU:      {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPATHMTU, Len: 4}},
		ssoPathMTU:             {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_PATHMTU, Len: sizeofIPv6Mtuinfo}},
		ssoChecksum:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_CHECKSUM, Len: 4}},
		ssoICMPFilter:          {Option: socket.Option{Level: iana.ProtocolIPv6ICMP, Name: sysICMP6_FILTER, Len: sizeofICMPv6Filter}},
		ssoJoinGroup:           {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:          {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		ssoTrafficClass:        {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_TCLASS, Len: 4}},
		ssoHopLimit:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_UNICAST_HOPS, Len: 4}},
		ssoMulticastInterface:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_IF, Len: 4}},
		ssoMulticastHopLimit:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_HOPS, Len: 4}},
		ssoMulticastLoopback:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTrafficClass: {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVTCLASS, Len: 4}},
		ssoReceiveHopLimit:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVHOPLIMIT, Len: 4}},
		ssoReceivePacketInfo:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPKTINFO, Len: 4}},
		ssoReceivePathMTU:      {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPATHMTU, Len: 4}},
		ssoPathMTU:             {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_PATHMTU, Len: sizeofIPv6Mtuinfo}},
		ssoChecksum:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_CHECKSUM, Len: 4}},
		ssoICMPFilter:          {Option: socket.Option{Level: iana.ProtocolIPv6ICMP, Name: sysICMP6_FILTER, Len: sizeofICMPv6Filter}},
		ssoJoinGroup:           {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:          {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
=======
		ssoTrafficClass:        {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_TCLASS, Len: 4}},
		ssoHopLimit:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_UNICAST_HOPS, Len: 4}},
		ssoMulticastInterface:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_IF, Len: 4}},
		ssoMulticastHopLimit:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_HOPS, Len: 4}},
		ssoMulticastLoopback:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTrafficClass: {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVTCLASS, Len: 4}},
		ssoReceiveHopLimit:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVHOPLIMIT, Len: 4}},
		ssoReceivePacketInfo:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVPKTINFO, Len: 4}},
		ssoReceivePathMTU:      {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_RECVPATHMTU, Len: 4}},
		ssoPathMTU:             {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_PATHMTU, Len: sizeofIPv6Mtuinfo}},
		ssoChecksum:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.IPV6_CHECKSUM, Len: 4}},
		ssoICMPFilter:          {Option: socket.Option{Level: iana.ProtocolIPv6ICMP, Name: unix.ICMP6_FILTER, Len: sizeofICMPv6Filter}},
		ssoJoinGroup:           {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:          {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: unix.MCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
)

var (
	ctlOpts = [ctlMax]ctlOpt{
		ctlTrafficClass: {sysIPV6_TCLASS, 4, marshalTrafficClass, parseTrafficClass},
		ctlHopLimit:     {sysIPV6_HOPLIMIT, 4, marshalHopLimit, parseHopLimit},
		ctlPacketInfo:   {sysIPV6_PKTINFO, sizeofInet6Pktinfo, marshalPacketInfo, parsePacketInfo},
		ctlNextHop:      {sysIPV6_NEXTHOP, sizeofSockaddrInet6, marshalNextHop, parseNextHop},
		ctlPathMTU:      {sysIPV6_PATHMTU, sizeofIPv6Mtuinfo, marshalPathMTU, parsePathMTU},
	}

	sockOpts = map[int]*sockOpt{
		ssoTrafficClass:        {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_TCLASS, Len: 4}},
		ssoHopLimit:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_UNICAST_HOPS, Len: 4}},
		ssoMulticastInterface:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_IF, Len: 4}},
		ssoMulticastHopLimit:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_HOPS, Len: 4}},
		ssoMulticastLoopback:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_MULTICAST_LOOP, Len: 4}},
		ssoReceiveTrafficClass: {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVTCLASS, Len: 4}},
		ssoReceiveHopLimit:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVHOPLIMIT, Len: 4}},
		ssoReceivePacketInfo:   {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPKTINFO, Len: 4}},
		ssoReceivePathMTU:      {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_RECVPATHMTU, Len: 4}},
		ssoPathMTU:             {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_PATHMTU, Len: sizeofIPv6Mtuinfo}},
		ssoChecksum:            {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysIPV6_CHECKSUM, Len: 4}},
		ssoICMPFilter:          {Option: socket.Option{Level: iana.ProtocolIPv6ICMP, Name: sysICMP6_FILTER, Len: sizeofICMPv6Filter}},
		ssoJoinGroup:           {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoLeaveGroup:          {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_GROUP, Len: sizeofGroupReq}, typ: ssoTypeGroupReq},
		ssoJoinSourceGroup:     {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_JOIN_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoLeaveSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_LEAVE_SOURCE_GROUP, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoBlockSourceGroup:    {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_BLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
		ssoUnblockSourceGroup:  {Option: socket.Option{Level: iana.ProtocolIPv6, Name: sysMCAST_UNBLOCK_SOURCE, Len: sizeofGroupSourceReq}, typ: ssoTypeGroupSourceReq},
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	}
)

func (sa *sockaddrInet6) setSockaddr(ip net.IP, i int) {
	sa.Family = syscall.AF_INET6
	copy(sa.Addr[:], ip)
	sa.Scope_id = uint32(i)
}

func (pi *inet6Pktinfo) setIfindex(i int) {
	pi.Ifindex = uint32(i)
}

func (mreq *ipv6Mreq) setIfindex(i int) {
	mreq.Interface = uint32(i)
}

func (gr *groupReq) setGroup(grp net.IP) {
	sa := (*sockaddrInet6)(unsafe.Pointer(uintptr(unsafe.Pointer(gr)) + 4))
	sa.Family = syscall.AF_INET6
	copy(sa.Addr[:], grp)
}

func (gsr *groupSourceReq) setSourceGroup(grp, src net.IP) {
	sa := (*sockaddrInet6)(unsafe.Pointer(uintptr(unsafe.Pointer(gsr)) + 4))
	sa.Family = syscall.AF_INET6
	copy(sa.Addr[:], grp)
	sa = (*sockaddrInet6)(unsafe.Pointer(uintptr(unsafe.Pointer(gsr)) + 260))
	sa.Family = syscall.AF_INET6
	copy(sa.Addr[:], src)
}
