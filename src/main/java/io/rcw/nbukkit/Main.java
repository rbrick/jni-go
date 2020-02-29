package io.rcw.nbukkit;

public class Main {

    public static void main(String[] args) {
        NativePlugin plugin = new NativePlugin("native-bukkit");
        System.out.println("Plugin name is " + plugin.getName());
        plugin.onEnable();


    }
}
