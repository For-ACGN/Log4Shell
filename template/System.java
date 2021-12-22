import java.io.*;

public class System {
    static {
        try {
            String[] command = {"${bin}", "${args}"};
            Process process = Runtime.getRuntime().exec(command);
            String line;

            InputStream inputStream = process.getInputStream();
            BufferedReader br = new BufferedReader(new InputStreamReader(inputStream));
            while((line = br.readLine()) != null){
                java.lang.System.out.println(line);
            }
        } catch (Exception e) {

        }
    }
}
