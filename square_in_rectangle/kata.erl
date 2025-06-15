-module(kata).

-export([squaresInRect/2]).

squaresInRect(Length, Width) when Length == Width ->
    [Length];
squaresInRect(Length, Width) ->
    if Length > Width ->
           [Width | squaresInRect(Length - Width, Width)];
       Width > Length ->
           [Length | squaresInRect(Length, Width)]
    end.
