#include <iostream>
#include <string>

using namespace std;

int main() {
    int cases;
    int rental, team;

    cin >> cases;

    while (cases--) {
        cin >> rental >> team;

        int price[rental];

        for (int i = 0; i < rental; i++) {
            cin >> price[i];
        }

        // 
        cout << rental << " - " << team << endl;
        for (int i = 0; i < rental; i++) {
            cout << price[i] << ", ";
        }
        cout << endl;
    }
}
