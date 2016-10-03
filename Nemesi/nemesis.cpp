using namespace std;

#include <list>
#include <vector>
#include <set>
#include <algorithm>

set<unsigned int> group(unsigned int kid, unsigned int kidsNum);
pair<set<unsigned int>, set<unsigned int>> removeEnemies(unsigned int kid, set<unsigned int> kidFriends, vector<unsigned int> nemesis);
void signalEnemies (unsigned int kid, set<unsigned int> enemies);

vector <unsigned int> v;
vector <unsigned int> kids = {};
vector<set<unsigned int>> groups = {};

void smista(int N, int nemico[]) {
	for (int i = 0; i < N; i++) {
		set<unsigned int> curGroup = group((unsigned int) i, (unsigned int) N);
		kids.insert(kids.begin(), (unsigned int) i);

		// Create group for i-th kid.
		groups.insert(groups.begin(), curGroup);
	}

	for (int i = 0; i < N; i++) {
		vector<unsigned int> kidEnemies = *(new vector<unsigned int>(nemico[i]));
		set<unsigned int> kidFriends = groups.at(i);
		pair<set<unsigned int>, set<unsigned int>> group = removeEnemies((unsigned int) i, kidFriends, kidEnemies);
		signalEnemies(i, group.second);
	}

    return;
}

set<unsigned int> group(unsigned int kid, unsigned int kidsNum) {
	set<unsigned int> kidGroup = *(new set<unsigned int>());

	for (unsigned int i = 0; i < kidsNum; i++)
		kidGroup.insert(i);

	return kidGroup;
}


pair<set<unsigned int>, set<unsigned int>> removeEnemies (unsigned int kid, set<unsigned int> kidFriends, vector<unsigned int> nemesis) {
	set<unsigned int> enemies = *(new set<unsigned int>());
	int m = nemesis.size();

	for (int i = 0; i < m; i++) {
		unsigned int enemy = nemesis.at(i);
		enemies.insert(enemy);
		kidFriends.erase(std::remove(kidFriends.begin(), kidFriends.end(), enemy), kidFriends.end());
	}


	return make_pair(kidFriends, enemies);
}

/* Signal kids that you are going to steal their goddam crayons. */
void signalEnemies (unsigned int kid, set<unsigned int> enemies) {
	std::set<unsigned int>::iterator i = enemies.begin();

	for (int k = 0; i != enemies.end(); i++, k++) {
		set<unsigned int> group = groups.at(k);
		group.erase(std::remove(group.begin(), group.end(), kid), group.end());
	}
}
