-module(reverse_word).

-export([reverse_word/1]).

reverse_word(Word) ->
    NewWord = string:split(Word, " "),
    TrimedList = lists:map(fun(W) -> string:trim(W) end, NewWord),
    string:join(
        lists:reverse(TrimedList), " ").
