euchredeck
==========

Euchre Baby!

A simple implementation for Euchre.

Well, not really, its more of a crude shuffler and dealer for Euchre sized decks and Euchre hands. It probably won't have an AI or a GUI, as I don't feel like writing a card gui in Go atm. There's better languages for that.

##Current Functionality
Currently creates a Euchre deck, shuffles it, and deals it.

The shuffle is an implementation of Fisher-Yates from github.com/mohae/shuffler

##Euchre basics
Euchre is a trick based, trump card game for four players played with a short deck, 24 cards. Players each form a side of a square with opposite sides oof the square being teammates. Games are played to 10 points.

###Euchre Deck and card hierarchy
A EuchreDeck is a special type of deck. It can either be viewed as half of a Pinochle Deck or a deck of 6 cards from each suit for a total of 24 cards: 9, 10, J, Q, K, A.

Card hierarchy is the same as traditional trump based card game hierarchy, except the lowest card is a 9: Ace is high, 9 is low. For trump, the Jack of the trump suit, or right bauer, and the Jack of the other suit of the same color as trump, the left bauer, are the highest and second highest card, respectively. So a trump of Clubs means that the Jack of Club is the right bauer, and the highest trump card, and the Jack of Spades is the left bauer, and the second highest trump card. Trump suit card order, from highest to lowest, is: right bauer, left bauer, Ace, King, Queen, 10, and 9.

##Determining the dealer
The dealer is traditionally determined by flipping each card face-up, in front of each player, going around in clock-wise order, until a Jack is dealt. That player deals the first hand.

###Dealing the cards
After shuffle, a cut is offerred to the player to the right of the dealer. The cut can either be accepted or declined. After the cut, the dealer deals each player 5 cards in two rounds, usually by 2 or 3 cards. The four remaining cards are placed on the table face down. The top-most card is turned face-up. The suit of this card is the first suit offered up for trump.

###Determining Trump
####Having the dealer 'pick it up'
Starting with the first player to the left of the dealer, each player has the option to tell the dealer to 'pick it up', which means that suit is trump and the caller's team will take 3 tricks in card play, or pass. If all the players pass, the dealer can either pick it up, or flip the card over, signifying that that suit has been rejected for trump.

If the dealer picks up the card, a card from his hand must be placed on the stack, face down. The dealer should have 5 cards in their hand after picking up the card.

####Calling Trump:
If the original suit was not called as trump by any of the players, a suit, other than the suit that was just rejected for trump, can be called by a player. The first player to have the option to call a suit is the player to the left of the dealer. This goes around until a suit has been called trump or the call becomes the dealer's choice.

Here, there are two options, or variations for Euchre. Some people play where the dealer must call trump. Other players play where the dealer can pass, at which point everybody returns their card and the deal passes on to the player to the left of the dealer.

##Playing the hand
Once trump has been determined, the first player to the left of the dealer leads the first trick by playing a card. Each playr to the left of that player plays their card, in order, until all 4 players have played a card. Players must follow suit, play a card with the same suit as the first card played, unless they are out of the suit. Then they have the option of either playing a trump card, if the suit led wasn't the trump suit, or play a  card from another suit, which is meaningless. 

The winner is the player who played the highest ranked card: which is either the highest ranked card of the suit led, or, the highest trump card if trump as played. The winner of the trick leads the next trick by playing a card first.

###Scoring
There are 5 tricks available in each hand. The team that called trump commits to winning at least three of those tricks. Doing so results in 1 point for the calling team; failing to do so results in 2 points for the opposing team--this is also called being 'euchred'. Winning all 5 tricks results in 2 points for the calling team.

If a player thinks they can win all 5 tricks alone, without the help of their partner, the player can say they are going alone when they either tell the dealer to 'pick it up' or call the trump. Going alone means that the partner lays down their cards, face down, and doesn't play. If the caller wins all 5 tricks, the team gets 4 points. If the caller wins 3 or 4 tricks, the team gets 1 point. Being Euchred results in 2 points, as usual.
