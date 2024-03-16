
class Solution {

    public long[] unmarkedSumArray(int[] nums, int[][] queries) {
        int len = nums.length;
        HashSet<Integer> ind = new HashSet<>();
        PriorityQueue<Integer> pq =new PriorityQueue<>();
        long sumOfNums= 0 ;
        for(int i=0;i<len;i++){
            ind.add(i);
            pq.add(nums[i]);
            sumOfNums+=nums[i];
        }
        int lent = queries.length;
        HashMap<Integer,Queue<Integer>> h = new HashMap<>();
        for(int i=0;i<len;i++){
            int val = nums[i];
            if(h.get(val)==null){
                Queue<Integer> q = new LinkedList<>();
                q.add(i);
                h.put(val,q);
            }else{
                Queue<Integer>  q= h.get(val);
                q.add(i);
            }
        }
        long ans[] = new long[lent];
        for(int i=0;i<lent;i++){
            int left = queries[i][0];
            int right = queries[i][1];
            
            long sum = 0;

            

            if(ind.contains(left)){

                ind.remove(left);

                sumOfNums-=nums[left];

            }

            

            while(right>0){

                

                if(pq.size()>0){

                    

                    int minval = pq.poll();

                    Queue<Integer> indexes = h.get(minval);

                    int minInd = indexes.remove();

                    

                    if(ind.contains(minInd)){

                        right--;

                        ind.remove(minInd);

                        sumOfNums-=nums[minInd];

                    }else{

                        

                    }

                    

                }else

                {

                    break;

                }

                 

            }

            ans[i] = sumOfNums;

        }

        

        return ans;

        

    }

} 