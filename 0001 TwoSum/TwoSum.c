typedef struct {
	int num;
	int idx;
	// Makes structure hashable
	UT_hash_handle hh;
} num_t;
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    // Brute Force
    // int* res = calloc((*returnSize = 2), sizeof(int));
    // for (int i = 0; i < numsSize - 1; i++) {
    //     for (int j = i + 1; j < numsSize; j++) {
    //         if (nums[i] + nums[j] == target) {
    //             res[0] = i;
    //             res[1] = j;
    //             break;
    //         }
    //     }
    // }
    // return res;

    // Two-pass Hash Table
    // int* res = calloc((*returnSize = 2), sizeof(int));
    // num_t *numsHash = NULL, *num = NULL, *tmp = NULL;
    // for (int i = 0; i < numsSize; i++) {
    //     num = malloc(sizeof(num_t));
    //     num->num = nums[i];
    //     num->idx = i;
    //     HASH_ADD_INT(numsHash, num, num);
    // }

    // for (int i = 0; i < numsSize - 1; i++) {
    //     int toFind = target - nums[i];
	// 	HASH_FIND_INT(numsHash, &toFind, num);

    //     if (num != NULL) {
    //         res[0] = i;
    //         res[1] = num->idx;
    //         break;
    //     }
    // }
    // return res;

    // One-pass Hash Table
    int* res = calloc((*returnSize = 2), sizeof(int));
	num_t *numsHash = NULL, *num = NULL, *tmp = NULL;

	for(int i = 0; i < numsSize; ++i){
		int toFind = target - nums[i];
		HASH_FIND_INT(numsHash, &toFind, num);

		if (num != NULL){
			res[0] = num->idx;
			res[1] = i;
			break;
		} else {
			num = malloc(sizeof(num_t));
			num->num = nums[i];
            num->idx = i;
			HASH_ADD_INT(numsHash, num, num);
		}
	}
	return res;
}