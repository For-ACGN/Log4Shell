public class Calc {
    static {
        try {
            Runtime.getRuntime().exec("calc");
        } catch (Exception e) {

        }
    }
}
