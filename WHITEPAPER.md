# SteveCare
This is a draft and notes of our whitepaper.  I am writing it while discussing with the community in our live stream on youtube. [Click here to join our livestream on youtube](https://www.youtube.com/watch?v=OrLvu8EZPj8).  You can also follow the project execution on our ugly-soon-to-be-beautiful website: [steve.care](https://steve.care) or on our sub-reddit: [/r/SteveCare](https://reddit.com/r/stevecare).

I am building a software that can fetch data from any source/protocol, including the standard (websites) and decentralized (bitcoin, ethereum, cosmos, defi, etc) web, emails, IRC, etc. Then make it easy to transform that data the way we want in order to present it in any format: visually on a screen, on file on disk, to a printer, over a network... anything we want!

I want non-technical people to be able to compose data with it in order to build their own software.  Therefore, I am also building a graphical user-interface (GUI) to make it easy to build software without code.

To make this possible, I want to create synergy between programmers and users.  Therefore, I will build a decentralized request system where non-technical users can request:

1. input/output combinations + complexity (the maximum amount of tokens reached during an execution) and/or
2. bad input that shouldn't work.  

That request will also contain a reward in units so that if a programmer sucessfully provide a grammar that matches the request the blockchain will transfer the units from the requester to the reward winners automatically.

When a programmer creates a script that matches a reward, he uses tokens that have been previously created by other programmers, or create new ones.  Those are basically NFT's (non-fungible tokens).  When a programmer wins a reward, that reward is separated between the token owners, the script writer and the creator of the SteveCare project as follow:

1. Each simple token (uint8 - 0-255 only values) being used directly costs 1 credit, given to the SteveCare project creator.
2. If a token executes 5 simple tokens on a script execution, the costs of that token would be 10 credits (5 to the creator, 5 to the token owner).
3. The script programmer basically submit a root token while answering a request, so he would get his revenue from the token he just registered.
4. If a reward of 200 units is paid for a request and the answer costs 2000 credits, that means that each credit owner of a reward receives 0.1 units for that reward.

The units will all be generated at the end of this project.  50% will be transfered to the people that helped promote the project while I was doing the live stream and for a period of 4 weeks following the live stream.  50% will be sold for bitcoin at the end of the live stream during a period of 4 weeks.

The amount of units given for each action on social media will be discussed with the community during the live stream. I want the community to be involved in this discussion so that it is fair for everyone.

The units given for bitcoin transactions won't involve humans at all, the project will read the bitcoin blockchain in order to reward depositors automatically.

I want to enable atomic swaps between bitcoin, dogecoin and monero against our units.  That way, people can interact with our decentralized software without any third party buying/selling units manually outside our ecosystem. This should reduce fraud attempts.

I also want people to be able to swap project owner shares and token NFT's against units. Therefore a programmer could sell his token NFT's (and its future revenue) for bitcoin, monero and/or dogecoin.

The amount of external coins to be swapped against our units will be expandable from our module system, by the community.

## Scopes
The project is separated in multiple scopes.  Here's the list of scopes.

### First Scope: Virtual Machine
The virtual machine is separated in four (4) sub-projects.  

1. The first one is called the lexer, its job is to make it possible to analyze data and tell us if input data matches a given grammar.  
2. The second one is called a parser, its job is to create a memory representation of input data using the returned tokens from the lexer.
3. The third one is called an interpreter, its job is to execute instructions based on the memory representation created by the parser application. The instructions can be core or module ones.
4. The fourth one is called modules, its job is to execute custom instructions in native code, so that the virtual machine is easily extendable by developers worldwide without having to modify the core engine of the virtual machine.

#### Lexer
A lexer is a software composed of a grammar.  The grammar contains a root token that is used to validate data.  The root token is composed of children tokens.

##### Token
1. The simplest form of a token is an uint8 (number between 0 and 255).  That number can represents ASCII characters or any kind of data.
2. Each token must have a cardinality
3. Each token can contain a list of children tokens.
4. Each element of the list can be an uint8 (0-255) or a child token.
5. Each token must be composed of lines of tokens.  If one (1) line is valid against data, it means that the token is valid.

##### Cardinality
Cardinality is the amount of times in which a token can be present in data. Its possible values are:

1. x, (x occurences or more)
2. Specific (x occurences)
3. Range (between x and y occurences)

#### Parser
To be determined

#### Interpreter
To be determined

#### Modules
To be determined

### First Scope: To be determined
To be determined
