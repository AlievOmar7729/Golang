#include <stdio.h>

double calculate_cost(int height, int width, int material, int glass, int window_sill) {
    double pricePerCm2 = 0.0;

    if (material == 0) { // Дерево
        if (glass == 0) pricePerCm2 = 2.5;
        else if (glass == 1) pricePerCm2 = 3.0; 
    }
    else if (material == 1) { // Металл
        if (glass == 0) pricePerCm2 = 0.5;
        else if (glass == 1) pricePerCm2 = 1.0;
    }
    else if (material == 2) { // Металлопластик
        if (glass == 0) pricePerCm2 = 1.5;
        else if (glass == 1) pricePerCm2 = 2.0;
    }

    double totalCost = (height * width) * pricePerCm2;

    if (window_sill) {
        totalCost += 350.0;
    }

    return totalCost;
}
