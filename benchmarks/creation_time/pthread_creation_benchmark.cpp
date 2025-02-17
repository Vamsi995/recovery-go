#include <iostream>
#include <pthread.h>
#include <vector>
#include <chrono>

using namespace std;
using namespace std::chrono;

// Thread function (does nothing)
void* emptyFunction(void*) { return nullptr; }

// Benchmarking pthread creation
void benchmarkPthreads(int num_threads) {
    vector<pthread_t> threads(num_threads);

    auto start = high_resolution_clock::now();

    for (int i = 0; i < num_threads; ++i) {
        pthread_create(&threads[i], nullptr, emptyFunction, nullptr);
    }

    for (int i = 0; i < num_threads; ++i) {
        pthread_join(threads[i], nullptr);
    }

    auto end = high_resolution_clock::now();
    duration<double, milli> elapsed = end - start;

    cout << "Created " << num_threads << " pthreads in " << elapsed.count() << " ms" << endl;
}

int main() {
    vector<int> testCases = {100, 1000, 10000, 100000}; // Avoid excessive threads

    cout << "Benchmarking pthread creation time:" << endl;
    for (int i = 0; i < testCases.size(); ++i) {
        benchmarkPthreads(testCases[i]);
    }

    return 0;
}

