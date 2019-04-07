

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QString>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>


class QmlBridged01c47: public QObject
{
Q_OBJECT
public:
	QmlBridged01c47(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlBridged01c47_QmlBridged01c47_QRegisterMetaType();QmlBridged01c47_QmlBridged01c47_QRegisterMetaTypes();callbackQmlBridged01c47_Constructor(this);};
	void Signal_SendToQml(QString data) { QByteArray ta17c9a = data.toUtf8(); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a.prepend("WHITESPACE").constData()+10), ta17c9a.size()-10 };callbackQmlBridged01c47_SendToQml(this, dataPacked); };
	 ~QmlBridged01c47() { callbackQmlBridged01c47_DestroyQmlBridge(this); };
	bool event(QEvent * e) { return callbackQmlBridged01c47_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlBridged01c47_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQmlBridged01c47_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlBridged01c47_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlBridged01c47_CustomEvent(this, event); };
	void deleteLater() { callbackQmlBridged01c47_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlBridged01c47_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlBridged01c47_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQmlBridged01c47_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlBridged01c47_TimerEvent(this, event); };
signals:
	void sendToQml(QString data);
public slots:
	void calculate(QString nStr) { QByteArray t5dbe96 = nStr.toUtf8(); Moc_PackedString nStrPacked = { const_cast<char*>(t5dbe96.prepend("WHITESPACE").constData()+10), t5dbe96.size()-10 };callbackQmlBridged01c47_Calculate(this, nStrPacked); };
private:
};

Q_DECLARE_METATYPE(QmlBridged01c47*)


void QmlBridged01c47_QmlBridged01c47_QRegisterMetaTypes() {
}

void QmlBridged01c47_ConnectSendToQml(void* ptr)
{
	QObject::connect(static_cast<QmlBridged01c47*>(ptr), static_cast<void (QmlBridged01c47::*)(QString)>(&QmlBridged01c47::sendToQml), static_cast<QmlBridged01c47*>(ptr), static_cast<void (QmlBridged01c47::*)(QString)>(&QmlBridged01c47::Signal_SendToQml));
}

void QmlBridged01c47_DisconnectSendToQml(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridged01c47*>(ptr), static_cast<void (QmlBridged01c47::*)(QString)>(&QmlBridged01c47::sendToQml), static_cast<QmlBridged01c47*>(ptr), static_cast<void (QmlBridged01c47::*)(QString)>(&QmlBridged01c47::Signal_SendToQml));
}

void QmlBridged01c47_SendToQml(void* ptr, struct Moc_PackedString data)
{
	static_cast<QmlBridged01c47*>(ptr)->sendToQml(QString::fromUtf8(data.data, data.len));
}

void QmlBridged01c47_Calculate(void* ptr, struct Moc_PackedString nStr)
{
	QMetaObject::invokeMethod(static_cast<QmlBridged01c47*>(ptr), "calculate", Q_ARG(QString, QString::fromUtf8(nStr.data, nStr.len)));
}

int QmlBridged01c47_QmlBridged01c47_QRegisterMetaType()
{
	return qRegisterMetaType<QmlBridged01c47*>();
}

int QmlBridged01c47_QmlBridged01c47_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlBridged01c47*>(const_cast<const char*>(typeName));
}

int QmlBridged01c47_QmlBridged01c47_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridged01c47>();
#else
	return 0;
#endif
}

int QmlBridged01c47_QmlBridged01c47_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridged01c47>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QmlBridged01c47___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QmlBridged01c47___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlBridged01c47___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QmlBridged01c47___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridged01c47___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridged01c47___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridged01c47___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridged01c47___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridged01c47___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridged01c47___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridged01c47___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridged01c47___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridged01c47___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridged01c47___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridged01c47___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QmlBridged01c47_NewQmlBridge(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridged01c47(static_cast<QWindow*>(parent));
	} else {
		return new QmlBridged01c47(static_cast<QObject*>(parent));
	}
}

void QmlBridged01c47_DestroyQmlBridge(void* ptr)
{
	static_cast<QmlBridged01c47*>(ptr)->~QmlBridged01c47();
}

void QmlBridged01c47_DestroyQmlBridgeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QmlBridged01c47_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlBridged01c47*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlBridged01c47_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlBridged01c47*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QmlBridged01c47_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridged01c47*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlBridged01c47_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridged01c47*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridged01c47_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridged01c47*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlBridged01c47_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlBridged01c47*>(ptr)->QObject::deleteLater();
}

void QmlBridged01c47_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridged01c47*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridged01c47_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridged01c47*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
