/*
  Copy of program here: https://alvinalexander.com/blog/post/java/java-program-consume-all-memory-ram-on-computer
  Run with: java -Xmx1024M MemoryEater
*/

import java.util.Vector;

public class MemoryEater
{
  public static void main(String[] args)
  {
    Vector v = new Vector();
    while (true)
    {
      byte b[] = new byte[1048576];
      v.add(b);
      Runtime rt = Runtime.getRuntime();
      System.out.println( "free memory: " + rt.freeMemory() );
    }
  }
}
