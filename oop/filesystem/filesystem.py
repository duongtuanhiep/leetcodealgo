from abc import ABC, abstractmethod
from typing import Callable


class MyFile(ABC):
    @abstractmethod
    def name(self) -> str:
        pass

    def size(self) -> int:
        return self.size

    def is_folder(self) -> bool:
        return False


class File(MyFile):
    def __init__(self):
        self.name = ""
        self.size = ""

    def is_folder() -> bool:
        return False


class Folder(MyFile):
    def __init__(self):
        self.name = ""
        self.size = ""
        self.child = []

    def is_folder() -> bool:
        return True

    pass


def openDir(directory) -> Folder:
    return Folder()


MatcherFunc = Callable[[MyFile], bool]


def matchName(name: str) -> MatcherFunc:
    def matchByName(m: MyFile) -> bool:
        return m.name() == name

    return matchByName


def match_size_range(upper, lower) -> MatcherFunc:
    # def match_by_size_range(m: MyFile) -> bool:
    #     return lower <= m.size() <= upper

    return lambda x: lower <= x.size() <= upper


def match_only_dir() -> MatcherFunc:
    return lambda x: x.is_folder()


def match_only_files() -> MatcherFunc:
    return lambda x: not x.is_folder()


# queries = [match_size_range(100,200), match_name("bar"), match_only_dir]
# finc("foo", queries)
def find(
    directory: MyFile,
    matchers: list[MatcherFunc],
    recursive=True,
) -> list[MyFile]:
    curDir = openDir(directory)

    res = []
    stack = [curDir]

    while stack:
        directory = stack.pop()
        for item in directory.child:
            if all(match(item) for match in matchers):
                stack.append(item)

            if recursive and item.isFolder():
                stack.append(item)

    return res


# if __name__ == "__main__":

#     def compareStr(first_str) -> callable:
#         def compare(second_str) -> bool:
#             return first_str == second_str

#         return compare

#     compareWithTommy = compareStr("tommy")

#     print(compareWithTommy("hifiger"))
