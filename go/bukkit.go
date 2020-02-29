package main

/*
 #include "jwrapper.h"
*/
import "C"

import (
	"fmt"
	"net/http"
)

/*
 * Class:     io_rcw_nbukkit_NativePlugin
 * Method:    onLoad
 * Signature: ()V
 */
//export Java_io_rcw_nbukkit_NativePlugin_onLoad
func Java_io_rcw_nbukkit_NativePlugin_onLoad(env *C.JNIEnv, clazz C.jobject) {
}

/*
 * Class:     io_rcw_nbukkit_NativePlugin
 * Method:    onEnable
 * Signature: ()V
 */
//export Java_io_rcw_nbukkit_NativePlugin_onEnable
func Java_io_rcw_nbukkit_NativePlugin_onEnable(env *C.JNIEnv, clazz C.jobject) {
	fmt.Println("Hello, World!")

	fmt.Println("starting http server...")

	http.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(wr, "This is running in Go from Java...WOWOOWOWOWOOW UWU")
	})

	http.ListenAndServe(":8080", nil)
}

/*
 * Class:     io_rcw_nbukkit_NativePlugin
 * Method:    onDisable
 * Signature: ()V
 */
//export Java_io_rcw_nbukkit_NativePlugin_onDisable
func Java_io_rcw_nbukkit_NativePlugin_onDisable(env *C.JNIEnv, clazz C.jobject) {}

/*
 * Class:     io_rcw_nbukkit_NativePlugin
 * Method:    getName
 * Signature: ()Ljava/lang/String;
 */
//export Java_io_rcw_nbukkit_NativePlugin_getName
func Java_io_rcw_nbukkit_NativePlugin_getName(env *C.JNIEnv, obj C.jobject) C.jstring {
	return C.createJString(env, C.CString("TotallyNotANativeGoPlugin"))
}
func main() {}
