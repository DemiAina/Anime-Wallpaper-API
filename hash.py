import os
import math

# Adding more file types later 
# jpg and png support
supported_file_types = ["jpg", "png"]
test_file = "wallhaven-2y3wr9.jpg" 
# time_of_upload = "19:34" // add later
# date = "2023-8-8"

def main():
    num_rep_types = convert_to_ascii(supported_file_types)
    jpg = num_rep_types[0]
    png = num_rep_types[1]
    count_before_type = count_in_file()
    file_before_hash  = test_file[:-4]
    file_after_hash = convert_to_hash(file_before_hash, count_before_type)
    complete_hash = (f"{jpg}{file_after_hash}{count_before_type}")
    print(complete_hash)

def convert_to_ascii(file_types: list) -> list: 
    converted_list = []
    temp = "" 
    for i in file_types:
        for x in i:
            x = ord(x)
            temp += str(x)
        converted_list.append(temp)
        temp = ""
    return converted_list

def count_in_file():
    count = 0
    for i in test_file:
        if i == ".":
            break
        else:
            count += 1
    return count

    

def convert_to_hash(file_to_hash: str, count: int) -> str: 
    temp = ""
    one_quarter = math.floor(count / 4)
    three_quarters = math.floor((count / 4) * 3)
    middle = math.floor(count / 2)
    
    one_quarter_str = file_to_hash[one_quarter]
    three_quarters_str = file_to_hash[three_quarters]
    middle_str = file_to_hash[middle]

    temp += one_quarter_str
    temp += three_quarters_str
    temp += middle_str

    return temp



if __name__ == "__main__":
    main()

