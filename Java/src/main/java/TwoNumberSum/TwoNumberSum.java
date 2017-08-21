package TwoNumberSum;

import java.util.Scanner;

/**
 * 问题: 给定两个长度相等正数，从高位到低位来得到SUM，这题出的真没意思
 * Created by paladintyrion on 17/8/21.
 *
 * @author <a href="mailto: paladinosmenttt@sina.com" /> paladintyrion
 * @version 1.0.0
 */
public class TwoNumberSum {

    public static void printSum(int len, String num1, String num2) {
        int mark = 0;
        int pre = 0;
        boolean start = false;
        for (int i = 0; i < len; i++) {
            int a = Integer.valueOf(num1.substring(i, i+1)) + Integer.valueOf(num2.substring(i, i+1));
            if (a < 9) {
                if (start) {
                    System.out.print(pre);
                }
                while (mark > 0) {
                    System.out.print(9);
                }
                pre = a;
                start = true;
            } else if (a == 9) {
                mark++;
            } else {
                System.out.print(pre+1);
                while (mark > 0) {
                    System.out.print(0);
                    mark--;
                }
                pre = a - 10;
                start = true;
            }
        }
        if (start) {
            System.out.print(pre);
        }
        while (mark > 0) {
            System.out.print(9);
            mark--;
        }
    }

    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        String[] num = new String[2];
        for (int i = 0; i < 2; i++){
            num[i] = scan.next();
        }
        if (num[0].length() != num[1].length()) {
            System.exit(0);
        }
        printSum(num[0].length(), num[0], num[1]);
    }
}
