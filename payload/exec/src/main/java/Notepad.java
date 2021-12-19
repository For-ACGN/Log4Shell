public class Notepad {
    static {
        try {
            Runtime.getRuntime().exec("notepad");
        } catch (Exception e) {

        }
    }
}
