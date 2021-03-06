//go:generate gsl sockopts.xml
package goczmq

/*  =========================================================================
    zsock_option - get/set 0MQ socket options

            ****************************************************
            *   GENERATED SOURCE CODE, DO NOT EDIT!!           *
            *   TO CHANGE THIS, EDIT sockopts.gsl              *
            *   AND RUN gsl -q sockopts.xml                    *
            ****************************************************

    Copyright (c) the Contributors as noted in the AUTHORS file.
    This file is part of goczmq, the high-level go binding for CZMQ:
    http://github.com/zeromq/goczmq

    This Source Code Form is subject to the terms of the Mozilla Public
    License, v. 2.0. If a copy of the MPL was not distributed with this
    file, You can obtain one at http://mozilla.org/MPL/2.0/.
    =========================================================================
*/

/*
#include "czmq.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"unsafe"
)

// SetTos sets the tos option for the socket
func (s *Sock) SetTos(val int) {
	C.zsock_set_tos(unsafe.Pointer(s.zsockT), C.int(val))
}

// Tos returns the current value of the socket's tos option
func (s *Sock) Tos() int {
	val := C.zsock_tos(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRouterHandover sets the router_handover option for the socket
func (s *Sock) SetRouterHandover(val int) {
	C.zsock_set_router_handover(unsafe.Pointer(s.zsockT), C.int(val))
}

// SetRouterMandatory sets the router_mandatory option for the socket
func (s *Sock) SetRouterMandatory(val int) {
	C.zsock_set_router_mandatory(unsafe.Pointer(s.zsockT), C.int(val))
}

// SetProbeRouter sets the probe_router option for the socket
func (s *Sock) SetProbeRouter(val int) {
	C.zsock_set_probe_router(unsafe.Pointer(s.zsockT), C.int(val))
}

// SetReqRelaxed sets the req_relaxed option for the socket
func (s *Sock) SetReqRelaxed(val int) {
	C.zsock_set_req_relaxed(unsafe.Pointer(s.zsockT), C.int(val))
}

// SetReqCorrelate sets the req_correlate option for the socket
func (s *Sock) SetReqCorrelate(val int) {
	C.zsock_set_req_correlate(unsafe.Pointer(s.zsockT), C.int(val))
}

// SetConflate sets the conflate option for the socket
func (s *Sock) SetConflate(val int) {
	C.zsock_set_conflate(unsafe.Pointer(s.zsockT), C.int(val))
}

// SetZapDomain sets the zap_domain option for the socket
func (s *Sock) SetZapDomain(val string) {
	C.zsock_set_zap_domain(unsafe.Pointer(s.zsockT), C.CString(val))
}

// ZapDomain returns the current value of the socket's zap_domain option
func (s *Sock) ZapDomain() string {
	val := C.zsock_zap_domain(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// Mechanism returns the current value of the socket's mechanism option
func (s *Sock) Mechanism() int {
	val := C.zsock_mechanism(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetPlainServer sets the plain_server option for the socket
func (s *Sock) SetPlainServer(val int) {
	C.zsock_set_plain_server(unsafe.Pointer(s.zsockT), C.int(val))
}

// PlainServer returns the current value of the socket's plain_server option
func (s *Sock) PlainServer() int {
	val := C.zsock_plain_server(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetPlainUsername sets the plain_username option for the socket
func (s *Sock) SetPlainUsername(val string) {
	C.zsock_set_plain_username(unsafe.Pointer(s.zsockT), C.CString(val))
}

// PlainUsername returns the current value of the socket's plain_username option
func (s *Sock) PlainUsername() string {
	val := C.zsock_plain_username(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetPlainPassword sets the plain_password option for the socket
func (s *Sock) SetPlainPassword(val string) {
	C.zsock_set_plain_password(unsafe.Pointer(s.zsockT), C.CString(val))
}

// PlainPassword returns the current value of the socket's plain_password option
func (s *Sock) PlainPassword() string {
	val := C.zsock_plain_password(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetCurveServer sets the curve_server option for the socket
func (s *Sock) SetCurveServer(val int) {
	C.zsock_set_curve_server(unsafe.Pointer(s.zsockT), C.int(val))
}

// CurveServer returns the current value of the socket's curve_server option
func (s *Sock) CurveServer() int {
	val := C.zsock_curve_server(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetCurvePublickey sets the curve_publickey option for the socket
func (s *Sock) SetCurvePublickey(val string) {
	C.zsock_set_curve_publickey(unsafe.Pointer(s.zsockT), C.CString(val))
}

// CurvePublickey returns the current value of the socket's curve_publickey option
func (s *Sock) CurvePublickey() string {
	val := C.zsock_curve_publickey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetCurveSecretkey sets the curve_secretkey option for the socket
func (s *Sock) SetCurveSecretkey(val string) {
	C.zsock_set_curve_secretkey(unsafe.Pointer(s.zsockT), C.CString(val))
}

// CurveSecretkey returns the current value of the socket's curve_secretkey option
func (s *Sock) CurveSecretkey() string {
	val := C.zsock_curve_secretkey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetCurveServerkey sets the curve_serverkey option for the socket
func (s *Sock) SetCurveServerkey(val string) {
	C.zsock_set_curve_serverkey(unsafe.Pointer(s.zsockT), C.CString(val))
}

// CurveServerkey returns the current value of the socket's curve_serverkey option
func (s *Sock) CurveServerkey() string {
	val := C.zsock_curve_serverkey(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetGssapiServer sets the gssapi_server option for the socket
func (s *Sock) SetGssapiServer(val int) {
	C.zsock_set_gssapi_server(unsafe.Pointer(s.zsockT), C.int(val))
}

// GssapiServer returns the current value of the socket's gssapi_server option
func (s *Sock) GssapiServer() int {
	val := C.zsock_gssapi_server(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetGssapiPlaintext sets the gssapi_plaintext option for the socket
func (s *Sock) SetGssapiPlaintext(val int) {
	C.zsock_set_gssapi_plaintext(unsafe.Pointer(s.zsockT), C.int(val))
}

// GssapiPlaintext returns the current value of the socket's gssapi_plaintext option
func (s *Sock) GssapiPlaintext() int {
	val := C.zsock_gssapi_plaintext(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetGssapiPrincipal sets the gssapi_principal option for the socket
func (s *Sock) SetGssapiPrincipal(val string) {
	C.zsock_set_gssapi_principal(unsafe.Pointer(s.zsockT), C.CString(val))
}

// GssapiPrincipal returns the current value of the socket's gssapi_principal option
func (s *Sock) GssapiPrincipal() string {
	val := C.zsock_gssapi_principal(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetGssapiServicePrincipal sets the gssapi_service_principal option for the socket
func (s *Sock) SetGssapiServicePrincipal(val string) {
	C.zsock_set_gssapi_service_principal(unsafe.Pointer(s.zsockT), C.CString(val))
}

// GssapiServicePrincipal returns the current value of the socket's gssapi_service_principal option
func (s *Sock) GssapiServicePrincipal() string {
	val := C.zsock_gssapi_service_principal(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetIpv6 sets the ipv6 option for the socket
func (s *Sock) SetIpv6(val int) {
	C.zsock_set_ipv6(unsafe.Pointer(s.zsockT), C.int(val))
}

// Ipv6 returns the current value of the socket's ipv6 option
func (s *Sock) Ipv6() int {
	val := C.zsock_ipv6(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetImmediate sets the immediate option for the socket
func (s *Sock) SetImmediate(val int) {
	C.zsock_set_immediate(unsafe.Pointer(s.zsockT), C.int(val))
}

// Immediate returns the current value of the socket's immediate option
func (s *Sock) Immediate() int {
	val := C.zsock_immediate(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRouterRaw sets the router_raw option for the socket
func (s *Sock) SetRouterRaw(val int) {
	C.zsock_set_router_raw(unsafe.Pointer(s.zsockT), C.int(val))
}

// SetIpv4only sets the ipv4only option for the socket
func (s *Sock) SetIpv4only(val int) {
	C.zsock_set_ipv4only(unsafe.Pointer(s.zsockT), C.int(val))
}

// Ipv4only returns the current value of the socket's ipv4only option
func (s *Sock) Ipv4only() int {
	val := C.zsock_ipv4only(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetDelayAttachOnConnect sets the delay_attach_on_connect option for the socket
func (s *Sock) SetDelayAttachOnConnect(val int) {
	C.zsock_set_delay_attach_on_connect(unsafe.Pointer(s.zsockT), C.int(val))
}

// Type returns the current value of the socket's type option
func (s *Sock) Type() int {
	val := C.zsock_type(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetSndhwm sets the sndhwm option for the socket
func (s *Sock) SetSndhwm(val int) {
	C.zsock_set_sndhwm(unsafe.Pointer(s.zsockT), C.int(val))
}

// Sndhwm returns the current value of the socket's sndhwm option
func (s *Sock) Sndhwm() int {
	val := C.zsock_sndhwm(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRcvhwm sets the rcvhwm option for the socket
func (s *Sock) SetRcvhwm(val int) {
	C.zsock_set_rcvhwm(unsafe.Pointer(s.zsockT), C.int(val))
}

// Rcvhwm returns the current value of the socket's rcvhwm option
func (s *Sock) Rcvhwm() int {
	val := C.zsock_rcvhwm(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetAffinity sets the affinity option for the socket
func (s *Sock) SetAffinity(val int) {
	C.zsock_set_affinity(unsafe.Pointer(s.zsockT), C.int(val))
}

// Affinity returns the current value of the socket's affinity option
func (s *Sock) Affinity() int {
	val := C.zsock_affinity(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetSubscribe sets the subscribe option for the socket
func (s *Sock) SetSubscribe(val string) {
	C.zsock_set_subscribe(unsafe.Pointer(s.zsockT), C.CString(val))
}

// SetUnsubscribe sets the unsubscribe option for the socket
func (s *Sock) SetUnsubscribe(val string) {
	C.zsock_set_unsubscribe(unsafe.Pointer(s.zsockT), C.CString(val))
}

// SetIdentity sets the identity option for the socket
func (s *Sock) SetIdentity(val string) {
	C.zsock_set_identity(unsafe.Pointer(s.zsockT), C.CString(val))
}

// Identity returns the current value of the socket's identity option
func (s *Sock) Identity() string {
	val := C.zsock_identity(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// SetRate sets the rate option for the socket
func (s *Sock) SetRate(val int) {
	C.zsock_set_rate(unsafe.Pointer(s.zsockT), C.int(val))
}

// Rate returns the current value of the socket's rate option
func (s *Sock) Rate() int {
	val := C.zsock_rate(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRecoveryIvl sets the recovery_ivl option for the socket
func (s *Sock) SetRecoveryIvl(val int) {
	C.zsock_set_recovery_ivl(unsafe.Pointer(s.zsockT), C.int(val))
}

// RecoveryIvl returns the current value of the socket's recovery_ivl option
func (s *Sock) RecoveryIvl() int {
	val := C.zsock_recovery_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetSndbuf sets the sndbuf option for the socket
func (s *Sock) SetSndbuf(val int) {
	C.zsock_set_sndbuf(unsafe.Pointer(s.zsockT), C.int(val))
}

// Sndbuf returns the current value of the socket's sndbuf option
func (s *Sock) Sndbuf() int {
	val := C.zsock_sndbuf(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRcvbuf sets the rcvbuf option for the socket
func (s *Sock) SetRcvbuf(val int) {
	C.zsock_set_rcvbuf(unsafe.Pointer(s.zsockT), C.int(val))
}

// Rcvbuf returns the current value of the socket's rcvbuf option
func (s *Sock) Rcvbuf() int {
	val := C.zsock_rcvbuf(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetLinger sets the linger option for the socket
func (s *Sock) SetLinger(val int) {
	C.zsock_set_linger(unsafe.Pointer(s.zsockT), C.int(val))
}

// Linger returns the current value of the socket's linger option
func (s *Sock) Linger() int {
	val := C.zsock_linger(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetReconnectIvl sets the reconnect_ivl option for the socket
func (s *Sock) SetReconnectIvl(val int) {
	C.zsock_set_reconnect_ivl(unsafe.Pointer(s.zsockT), C.int(val))
}

// ReconnectIvl returns the current value of the socket's reconnect_ivl option
func (s *Sock) ReconnectIvl() int {
	val := C.zsock_reconnect_ivl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetReconnectIvlMax sets the reconnect_ivl_max option for the socket
func (s *Sock) SetReconnectIvlMax(val int) {
	C.zsock_set_reconnect_ivl_max(unsafe.Pointer(s.zsockT), C.int(val))
}

// ReconnectIvlMax returns the current value of the socket's reconnect_ivl_max option
func (s *Sock) ReconnectIvlMax() int {
	val := C.zsock_reconnect_ivl_max(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetBacklog sets the backlog option for the socket
func (s *Sock) SetBacklog(val int) {
	C.zsock_set_backlog(unsafe.Pointer(s.zsockT), C.int(val))
}

// Backlog returns the current value of the socket's backlog option
func (s *Sock) Backlog() int {
	val := C.zsock_backlog(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetMaxmsgsize sets the maxmsgsize option for the socket
func (s *Sock) SetMaxmsgsize(val int) {
	C.zsock_set_maxmsgsize(unsafe.Pointer(s.zsockT), C.int(val))
}

// Maxmsgsize returns the current value of the socket's maxmsgsize option
func (s *Sock) Maxmsgsize() int {
	val := C.zsock_maxmsgsize(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetMulticastHops sets the multicast_hops option for the socket
func (s *Sock) SetMulticastHops(val int) {
	C.zsock_set_multicast_hops(unsafe.Pointer(s.zsockT), C.int(val))
}

// MulticastHops returns the current value of the socket's multicast_hops option
func (s *Sock) MulticastHops() int {
	val := C.zsock_multicast_hops(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetRcvtimeo sets the rcvtimeo option for the socket
func (s *Sock) SetRcvtimeo(val int) {
	C.zsock_set_rcvtimeo(unsafe.Pointer(s.zsockT), C.int(val))
}

// Rcvtimeo returns the current value of the socket's rcvtimeo option
func (s *Sock) Rcvtimeo() int {
	val := C.zsock_rcvtimeo(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetSndtimeo sets the sndtimeo option for the socket
func (s *Sock) SetSndtimeo(val int) {
	C.zsock_set_sndtimeo(unsafe.Pointer(s.zsockT), C.int(val))
}

// Sndtimeo returns the current value of the socket's sndtimeo option
func (s *Sock) Sndtimeo() int {
	val := C.zsock_sndtimeo(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetXpubVerbose sets the xpub_verbose option for the socket
func (s *Sock) SetXpubVerbose(val int) {
	C.zsock_set_xpub_verbose(unsafe.Pointer(s.zsockT), C.int(val))
}

// SetTcpKeepalive sets the tcp_keepalive option for the socket
func (s *Sock) SetTcpKeepalive(val int) {
	C.zsock_set_tcp_keepalive(unsafe.Pointer(s.zsockT), C.int(val))
}

// TcpKeepalive returns the current value of the socket's tcp_keepalive option
func (s *Sock) TcpKeepalive() int {
	val := C.zsock_tcp_keepalive(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetTcpKeepaliveIdle sets the tcp_keepalive_idle option for the socket
func (s *Sock) SetTcpKeepaliveIdle(val int) {
	C.zsock_set_tcp_keepalive_idle(unsafe.Pointer(s.zsockT), C.int(val))
}

// TcpKeepaliveIdle returns the current value of the socket's tcp_keepalive_idle option
func (s *Sock) TcpKeepaliveIdle() int {
	val := C.zsock_tcp_keepalive_idle(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetTcpKeepaliveCnt sets the tcp_keepalive_cnt option for the socket
func (s *Sock) SetTcpKeepaliveCnt(val int) {
	C.zsock_set_tcp_keepalive_cnt(unsafe.Pointer(s.zsockT), C.int(val))
}

// TcpKeepaliveCnt returns the current value of the socket's tcp_keepalive_cnt option
func (s *Sock) TcpKeepaliveCnt() int {
	val := C.zsock_tcp_keepalive_cnt(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetTcpKeepaliveIntvl sets the tcp_keepalive_intvl option for the socket
func (s *Sock) SetTcpKeepaliveIntvl(val int) {
	C.zsock_set_tcp_keepalive_intvl(unsafe.Pointer(s.zsockT), C.int(val))
}

// TcpKeepaliveIntvl returns the current value of the socket's tcp_keepalive_intvl option
func (s *Sock) TcpKeepaliveIntvl() int {
	val := C.zsock_tcp_keepalive_intvl(unsafe.Pointer(s.zsockT))
	return int(val)
}

// SetTcpAcceptFilter sets the tcp_accept_filter option for the socket
func (s *Sock) SetTcpAcceptFilter(val string) {
	C.zsock_set_tcp_accept_filter(unsafe.Pointer(s.zsockT), C.CString(val))
}

// TcpAcceptFilter returns the current value of the socket's tcp_accept_filter option
func (s *Sock) TcpAcceptFilter() string {
	val := C.zsock_tcp_accept_filter(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

// Rcvmore returns the current value of the socket's rcvmore option
func (s *Sock) Rcvmore() int {
	val := C.zsock_rcvmore(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Fd returns the current value of the socket's fd option
func (s *Sock) Fd() int {
	val := C.zsock_fd(unsafe.Pointer(s.zsockT))
	return int(val)
}

// Events returns the current value of the socket's events option
func (s *Sock) Events() int {
	val := C.zsock_events(unsafe.Pointer(s.zsockT))
	return int(val)
}

// LastEndpoint returns the current value of the socket's last_endpoint option
func (s *Sock) LastEndpoint() string {
	val := C.zsock_last_endpoint(unsafe.Pointer(s.zsockT))
	return C.GoString(val)
}

