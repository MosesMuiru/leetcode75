-module(reverse).

-export([count_elements/1, get_index/2, convert_to_list/1, check_if_vowel_exists/1]).

%%-type element_t() :: integer() | string().

count_elements(List) ->
    count_elements(0, List).

count_elements(Counter, []) ->
    Counter;
count_elements(Counter, [_H | T]) ->
    count_elements(Counter + 1, T).

get_index(Element, List) ->
    get_index(Element, 0, [], List).

get_index(_Element, _Counter, Index, []) ->
    Index;
get_index(Element, Counter, Index, [H | T]) ->
    case H of
        Element ->
            get_index(Element, Counter + 1, Index ++ [Counter], T);
        _ ->
            get_index(Element, Counter + 1, Index, T)
    end.

%% then we can go to the
%%
convert_to_list(A_String) ->
    lists:map(fun(C) -> [C] end, A_String).

check_if_vowel_exists(Word) ->
    Vowels = ["A", "E", "I", "O", "U"],
    WordList = convert_to_list(Word),

    Index_of_letters =
        lists:map(fun(Letter) ->
                     % check if the the words contains a vowel
                     Upper_letter = string:to_upper(Letter),
                     lists:map(fun(Vowel) ->
                                  if Upper_letter == Vowel -> get_index(Letter, WordList) end
                               end,
                               Vowels)
                  end,
                  WordList),

    Index_of_letters.
