package main

/*
#cgo CFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -W -Wextra -ffunction-sections -fdata-sections -DUNICODE -D_UNICODE -DWIN32 -DMINGW_HAS_SECURE_API=1 -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_MULTIMEDIA_LIB -DQT_WIDGETS_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_DBUS_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -fno-keep-inline-dllexport -O2 -std=gnu++11 -Wall -W -Wextra -ffunction-sections -fdata-sections -fexceptions -mthreads -DUNICODE -D_UNICODE -DWIN32 -DMINGW_HAS_SECURE_API=1 -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_MULTIMEDIA_LIB -DQT_WIDGETS_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_NETWORK_LIB -DQT_DBUS_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../qt -I. -I/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/include -I/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/include/QtMultimedia -I/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/include/QtWidgets -I/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/include/QtQuick -I/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/include/QtGui -I/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/include/QtQml -I/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/include/QtNetwork -I/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/include/QtDBus -I/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/include/QtCore -Irelease -I/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/mkspecs/win32-g++
#cgo LDFLAGS:        -Wl,-subsystem,windows -Wl,--gc-sections -mthreads
#cgo LDFLAGS:        -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/mediaservice /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/mediaservice/libdsengine.a -lmf -lmfplat -lmfuuid -ld3d9 -ldxva2 -levr -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/lib -ldmoguids -lmsdmo -lksuser /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/mediaservice/libqtmedia_audioengine.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/audio /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/audio/libqtaudio_windows.a -lstrmiids -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/playlistformats /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/playlistformats/libqtmultimedia_m3u.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/styles /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/styles/libqwindowsvistastyle.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/platforms /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/platforms/libqwindows.a -lwinspool -lwtsapi32 /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5EventDispatcherSupport.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5FontDatabaseSupport.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5ThemeSupport.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5AccessibilitySupport.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5WindowsUIAutomationSupport.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/imageformats /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/imageformats/libqgif.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/imageformats/libqicns.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/imageformats/libqico.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/imageformats/libqjp2.a -ljasper /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/imageformats/libqjpeg.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/imageformats/libqmng.a -lmng -llcms2 -lpthread /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/imageformats/libqtga.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/imageformats/libqtiff.a -ltiff -llzma -ljpeg /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/imageformats/libqwbmp.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/imageformats/libqwebp.a -lwebpdemux -lwebp -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling/libqmldbg_debugger.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling/libqmldbg_inspector.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling/libqmldbg_local.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling/libqmldbg_messages.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling/libqmldbg_native.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling/libqmldbg_nativedebugger.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling/libqmldbg_preview.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling/libqmldbg_profiler.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling/libqmldbg_quickprofiler.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling/libqmldbg_server.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5PacketProtocol.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/qmltooling/libqmldbg_tcp.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/bearer /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/plugins/bearer/libqgenericbearer.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5Multimedia.a -lamstrmid /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5Widgets.a -luxtheme -ldwmapi /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5DBus.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/lib -ldbus-1 -ldbghelp -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick.2 /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick.2/libqtquick2plugin.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtGraphicalEffects/private /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtGraphicalEffects/private/libqtgraphicaleffectsprivate.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Window.2 /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Window.2/libwindowplugin.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtGraphicalEffects /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtGraphicalEffects/libqtgraphicaleffectsplugin.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Templates.2 /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Templates.2/libqtquicktemplates2plugin.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Controls.2 /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Controls.2/libqtquickcontrols2plugin.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Controls.2/Fusion /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Controls.2/Fusion/libqtquickcontrols2fusionstyleplugin.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Controls.2/Imagine /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Controls.2/Imagine/libqtquickcontrols2imaginestyleplugin.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Controls.2/Material /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Controls.2/Material/libqtquickcontrols2materialstyleplugin.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Controls.2/Universal /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Controls.2/Universal/libqtquickcontrols2universalstyleplugin.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5QuickControls2.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5QuickTemplates2.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Layouts /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/QtQuick/Layouts/libqquicklayoutsplugin.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5Quick.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5Gui.a -lharfbuzz -lcairo -lgobject-2.0 -lfontconfig -lfreetype -lm -lusp10 -lmsimg32 -lpixman-1 -lffi -lexpat -lbz2 -lpng16 -lharfbuzz_too -lfreetype_too -lglib-2.0 -lshlwapi -lpcre -lintl -liconv -lcomdlg32 -loleaut32 -limm32 -lopengl32 -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/Qt/labs/settings /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/qml/Qt/labs/settings/libqmlsettingsplugin.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5Qml.a /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5Network.a -ldnsapi -liphlpapi -lssl -lcrypto -lgdi32 -lcrypt32 /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5Core.a -lmpr -lnetapi32 -luserenv -lversion -lws2_32 -lkernel32 -luser32 -lshell32 -luuid -lole32 -ladvapi32 -lwinmm -lz -lpcre2-16  -lmingw32 /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libqtmain.a -L/usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib /usr/lib/mxe/usr/x86_64-w64-mingw32.static/qt5/lib/libQt5Core.a -lmpr -lnetapi32 -luserenv -lversion -lws2_32 -lkernel32 -luser32 -lshell32 -luuid -lole32 -ladvapi32 -lwinmm -lz -lpcre2-16 -lopengl32
#cgo LDFLAGS: -Wl,--allow-multiple-definition
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"