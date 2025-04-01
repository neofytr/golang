#include <bits/stdc++.h>

using namespace std;

#include <iostream>
#include <vector>

using namespace std;

void getCombinations(int id, int target, vector<int> &coins, vector<int> &path)
{
    if (!target)
    {
        for (int coin : path)
        {
            cout << coin << " ";
        }
        cout << endl;
        return;
    }

    if (target < 0 || id >= coins.size())
    {
        return;
    }

    path.push_back(coins[id]);
    getCombinations(id, target - coins[id], coins, path);
    path.pop_back();

    getCombinations(id + 1, target, coins, path);
}

int main()
{
    vector<int> coins = {2};
    int target = 3;
    vector<int> path;

    getCombinations(0, target, coins, path);
    return 0;
}