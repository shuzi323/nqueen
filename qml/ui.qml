import QtQuick 2.2
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3


Rectangle {
    visible: true
    property int margin: 11
    width: 400
    height: 460
    //minimumWidth: 600
    //minimumHeight: 450

    Image {                     //设置背景图片
        anchors.fill: parent
        id: myimage
        opacity: 0.4            //设置透明度
        fillMode: Image.PreserveAspectCrop
        source: "image/background.png"
    }

    ColumnLayout {
        id: mainLayout
        anchors.fill: parent
        anchors.margins: margin

        GroupBox {
            id: rowBox
            title: "input"
            Layout.fillWidth: true
            Layout.fillHeight: true
            //anchors.top: parent.topMargin
            opacity: 0.6

            RowLayout {
                id: rowLayout
                anchors.fill: parent

                TextField {
                    id: myInput
                    selectByMouse: true
                    focus: true
                    Keys.onReturnPressed: checkAndSend(myInput.text)  //按下enter键确认
                    validator: IntValidator{bottom: 0; top: 20;}
                    placeholderText: "Enter a number between 0 and 18."
                    Layout.fillWidth: true

                    function checkAndSend(n){
                        if (n !== "" && n>=0){
                            QmlBridge.calculate(String(n));
                        }
                    }
                }

                Button {
                    text: "Button"
                    onClicked: QmlBridge.calculate(String(myInput.text))
                }
            }
        }

        Flickable {
                  Layout.minimumHeight: parent.height/4*3
                  Layout.fillWidth: true
                  Layout.fillHeight: true
                  clip: true  //裁切掉超出区域的部分

                  contentHeight: t3.height
                  contentWidth: parent.width

                  TextArea {
                    id: t3
                    anchors.top: parent.top
                    readOnly: true
                    selectByMouse: true
                    text: ""
                    //Layout.minimumHeight: parent.height/4*3 
                    //Layout.fillHeight: true
                    //Layout.fillWidth: true
                    
                    Connections{
                        target: QmlBridge
                        onSendToQml: t3.text = data + "\n\n" + t3.text
                        }
                    }
                  
                  ScrollBar.vertical: ScrollBar { }
              }
    }
}