package io.rcw.nbukkit;

public final class NativePlugin {

    public NativePlugin(String libname) {
        System.loadLibrary(libname);
    }

    native void onLoad();

    native void onEnable();

    native void onDisable();

    native String getName();
}
