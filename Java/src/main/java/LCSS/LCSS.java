package LCSS;

/**
 * Created by paladintyrion on 17/7/21.
 *
 * @author <a href="mailto: paladinosmenttt@sina.com" /> paladintyrion
 * @version 1.0.0
 */

/**
 * 最大连续子序列和问题：
 * 在一列数中寻找一些数，这些数满足：
 * 子序列中，∑()最大。
 * 需要明确的是状态转移方程中的状态代表的含义。因为contiguous，
 * 所以dp[i]代表的应该以i位置元素结尾的连续值，并非最大值。
 */
public class LCSS {

    public static int getMaxSubArray(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }

        int maxEndingHere = 0;
        int trueMax = Integer.MIN_VALUE;
        for (int i = 0; i < nums.length; i++) {
            if (maxEndingHere < 0) {
                maxEndingHere = 0;
            }
            maxEndingHere += nums[i];
            trueMax = Math.max(trueMax, maxEndingHere);
            System.out.println(trueMax);
        }
        return trueMax;
    }

    public static void main(String[] args) {
        int[] s = new int[6];
        s[0] = 99;
        s[1] = -7;
        s[2] = -8;
        s[3] = 8;
        s[4] = 88;
        s[5] = 3;

        int res = getMaxSubArray(s);
        System.out.println(res);
    }
}
