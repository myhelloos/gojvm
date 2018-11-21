package jvmgo.book;

public class GetClassTest {
  public static void main(String[] args) {
      System.out.println(void.class.getName()); // void
      System.out.println(boolean.class.getName()); // void
      System.out.println(byte.class.getName()); // void
      System.out.println(char.class.getName()); // void
      System.out.println(short.class.getName()); // void
      System.out.println(int.class.getName()); // void
      System.out.println(long.class.getName()); // void
      System.out.println(float.class.getName()); // void
      System.out.println(double.class.getName()); // void
      System.out.println(Object.class.getName()); // void
      System.out.println(int[].class.getName()); // void
      System.out.println(int[][].class.getName()); // void
      System.out.println(Object[].class.getName()); // void
      System.out.println(Object[][].class.getName()); // void
      System.out.println(Runnable.class.getName()); // void
      System.out.println("abc".getClass().getName()); // void
      System.out.println(new double[0].getClass().getName()); // void
      System.out.println(new String[0].getClass().getName()); // void
  }
}
