from collections import deque


def search(name):
    search_queue = deque()
    search_queue += graph[name]
    searched = []

    while search_queue:
        person = search_queue.popleft()
        if not person in searched:
            if isFirst(person):
                print(person)
                return True
            else:
                search_queue += graph[person]
                searched.append(person)
    return False


def isFirst(name):
    return name == "one"


graph = {
    "one": ["two", "three"],
    "two": ["three"],
    "three": ["one"],
}

search("one")
