package jvmgo.book;

public class ArrayDemo {
  public static void main(String[] args) {
      int[] a1 = new int[10];
      String[] a2 = new String[10];
      int[][] a3 = new int[10][10];
      int x = a1.length;
      a1[0] = 100;
      int y = a1[0];
      a2[0] = "abc";
      String s = a2[0];
  }
}
