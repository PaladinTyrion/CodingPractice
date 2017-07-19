package LIS;

/**
 * Created by paladintyrion on 17/7/19.
 *
 * @author <a href="mailto: paladinosmenttt@sina.com" /> paladintyrion
 * @version 1.0.0
 */

import java.util.ArrayList;
import java.util.Collections;

/**
 * 最大上升子序列问题
 * 在一列数中寻找一些数，这些数满足：
 * 任意两个数a[i]和a[j]，若i<j，必有a[i]<a[j]，这样最长的子序列称为最长递增（上升）子序列。
 *
 * (1)
 * 设dp[i]表示以i为结尾的最长递增子序列的长度，则状态转移方程为：
 * dp[i] = max{dp[j]+1}, 1<=j<i,a[j]<a[i].时间复杂度为O(n*n)；
 *
 * (2)
 * 考虑两个数a[x]和a[y]，x<y且a[x]<a[y],且dp[x]=dp[y]，那么我们该选择哪个呢？
 * 显然a[x],因为它更有潜力，也就是说我们可以用a[x]来替换掉a[y]，
 * 也就是说我们需要维护一个数据结构来存储可能的递增子序列的元素，并且需要在某些时候进行替换。
 * 因此我们可以用一个链表来存储，并且在查找替换位置的时候用二分查找来实现，这样时间复杂度为O(nlogn)。
 */
public class LIS {

    /**
     * 上述中(1)的求解方法,时间复杂度为O(n^2)
     * @param nums
     * @return
     */
    public static int getLis(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        int max = 1;
        int[] dp = new int[nums.length];
        for (int i = 0; i < nums.length; i++) {
            dp[i] = 1;
            for (int j = 0; j < i; j++) {
                if (nums[j] < nums[i]) {
                    dp[i] = Math.max(dp[i], dp[j] + 1);
                    max = Math.max(max, dp[i]);
                }
            }
        }
        return max;
    }

    /**
     * 第二种求节方法，可以直接返回dp.size()也可以得到最大上升子序列长度，时间复杂度O(nlogn)
     * @param nums
     * @return
     */
    public static ArrayList<Integer> getLIS(int[] nums) {
        if (nums == null || nums.length == 0) {
            return null;
        }
        ArrayList<Integer> dp = new ArrayList<>();
        for (int item : nums) {
            if (dp.size() == 0 || dp.get(dp.size() - 1) < item) {
                dp.add(item);
                System.out.println("1: " + dp);
            } else {
                int i = Collections.binarySearch(dp, item);//insert position
                dp.set(i < 0 ? -i - 1 : i, item);
                System.out.println("2: " + dp);
            }
        }
        return dp;
    }

    //输出为6，最大上升子序列为4 6 11 13 18 51
    public static void main(String[] args) {
        int[] nums = new int[9];
        nums[0] = 8;
        nums[1] = 4;
        nums[2] = 6;
        nums[3] = 22;
        nums[4] = 11;
        nums[5] = 13;
        nums[6] = 29;
        nums[7] = 18;
        nums[8] = 51;

        int res = getLis(nums);
        System.out.println(res);


        ArrayList<Integer> dp = getLIS(nums);
        System.out.println(dp);
    }
}
