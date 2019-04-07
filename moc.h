

#pragma once

#ifndef GO_MOC_d01c47_H
#define GO_MOC_d01c47_H

#include <stdint.h>

#ifdef __cplusplus
class QmlBridged01c47;
void QmlBridged01c47_QmlBridged01c47_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void QmlBridged01c47_ConnectSendToQml(void* ptr);
void QmlBridged01c47_DisconnectSendToQml(void* ptr);
void QmlBridged01c47_SendToQml(void* ptr, struct Moc_PackedString data);
void QmlBridged01c47_Calculate(void* ptr, struct Moc_PackedString nStr);
int QmlBridged01c47_QmlBridged01c47_QRegisterMetaType();
int QmlBridged01c47_QmlBridged01c47_QRegisterMetaType2(char* typeName);
int QmlBridged01c47_QmlBridged01c47_QmlRegisterType();
int QmlBridged01c47_QmlBridged01c47_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlBridged01c47___dynamicPropertyNames_atList(void* ptr, int i);
void QmlBridged01c47___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlBridged01c47___dynamicPropertyNames_newList(void* ptr);
void* QmlBridged01c47___findChildren_atList2(void* ptr, int i);
void QmlBridged01c47___findChildren_setList2(void* ptr, void* i);
void* QmlBridged01c47___findChildren_newList2(void* ptr);
void* QmlBridged01c47___findChildren_atList3(void* ptr, int i);
void QmlBridged01c47___findChildren_setList3(void* ptr, void* i);
void* QmlBridged01c47___findChildren_newList3(void* ptr);
void* QmlBridged01c47___findChildren_atList(void* ptr, int i);
void QmlBridged01c47___findChildren_setList(void* ptr, void* i);
void* QmlBridged01c47___findChildren_newList(void* ptr);
void* QmlBridged01c47___children_atList(void* ptr, int i);
void QmlBridged01c47___children_setList(void* ptr, void* i);
void* QmlBridged01c47___children_newList(void* ptr);
void* QmlBridged01c47_NewQmlBridge(void* parent);
void QmlBridged01c47_DestroyQmlBridge(void* ptr);
void QmlBridged01c47_DestroyQmlBridgeDefault(void* ptr);
char QmlBridged01c47_EventDefault(void* ptr, void* e);
char QmlBridged01c47_EventFilterDefault(void* ptr, void* watched, void* event);
void QmlBridged01c47_ChildEventDefault(void* ptr, void* event);
void QmlBridged01c47_ConnectNotifyDefault(void* ptr, void* sign);
void QmlBridged01c47_CustomEventDefault(void* ptr, void* event);
void QmlBridged01c47_DeleteLaterDefault(void* ptr);
void QmlBridged01c47_DisconnectNotifyDefault(void* ptr, void* sign);
void QmlBridged01c47_TimerEventDefault(void* ptr, void* event);
;

#ifdef __cplusplus
}
#endif

#endif