def dailyTemperatures(temperatures):
    n = len(temperatures)
    answer = [0] * n
    pile = []

    for i in range(n):
        while pile and temperatures[i] > temperatures[pile[-1]]:
            prev_index = pile.pop()
            answer[prev_index] = i - prev_index
        pile.append(i)

    return answer
