## SLRZ - Subtraction (Left | Right) Zone

A combinatorial game theory analysis of a partizan subtraction game with zone-based multipliers.

---

## Table of Contents

- [Introduction](#introduction)
- [Game Rules](#game-rules)
- [Sample Game](#sample-game)
- [Outcome Classes](#outcome-classes)
- [Analysis Results](#analysis-results)
- [Key Findings](#key-findings)
- [The Smaller Option Advantage](#the-smaller-option-advantage)

---

## Introduction

**SLRZ (Subtraction Left | Right Zone)** is a partizan combinatorial game and a variant of Subtraction (L | R). Two players, **Left** and **Right**, take turns subtracting from a stack of `n` counters. Unlike traditional subtraction games, SLRZ introduces a **zone system** that can double a player's subtraction under certain conditions.

This project analyzes 10,000 game positions where:
- Left's options (L) ∈ {1, 2, 3, ..., 10}
- Right's options (R) ∈ {1, 2, 3, ..., 10}  
- Counter values (n) ∈ {1, 2, 3, ..., 100}

---

## Game Rules

### Basic Mechanics

1. **Players**: Two players called **Left (L)** and **Right (R)**
2. **Starting Position**: A stack of `n` counters
3. **Moves**: 
   - Left subtracts exactly `L` counters on their turn
   - Right subtracts exactly `R` counters on their turn
4. **Losing Condition**: A player who cannot make a legal move (would result in negative counters) **loses**

### The Zone System

The counters exist in different **zones** based on their value modulo 5:

| n mod 5 | Zone | Owner |
|---------|------|-------|
| 0 | Left's Zone (LZ) | Left |
| 1 | Left's Zone (LZ) | Left |
| 2 | Neutral | Neither |
| 3 | Right's Zone (RZ) | Right |
| 4 | Right's Zone (RZ) | Right |

### Zone Multiplier Rule

When a player's move **lands in the opponent's zone**, their subtraction is **doubled**:

- If **Left** subtracts `L` and the result `(n - L)` is in **Right's Zone** (RZ), Left instead subtracts `2L`
- If **Right** subtracts `R` and the result `(n - R)` is in **Left's Zone** (LZ), Right instead subtracts `2R`

**Exception**: The multiplier does NOT apply if the subtraction would result in exactly 0 counters.

---

## Sample Game

**Setup**: L = 3, R = 5, n = 17 (starting counters)

```
Starting Position: 17 counters
Left's option: 3
Right's option: 5
```

### Turn-by-Turn Analysis

**Turn 1 - Left moves:**
```
Position: 17 (17 mod 5 = 2, Neutral zone)
Left wants to subtract 3: 17 - 3 = 14
Check: 14 mod 5 = 4 → Right's Zone (RZ)!
DOUBLED: Left subtracts 6 instead: 17 - 6 = 11
```

**Turn 2 - Right moves:**
```
Position: 11 (11 mod 5 = 1, Left's Zone)
Right wants to subtract 5: 11 - 5 = 6
Check: 6 mod 5 = 1 → Left's Zone (LZ)!
DOUBLED: Right subtracts 10 instead: 11 - 10 = 1
```

**Turn 3 - Left moves:**
```
Position: 1 (1 mod 5 = 1, Left's Zone)
Left wants to subtract 3: 1 - 3 = -2
ILLEGAL MOVE: Result is negative!
Left cannot move → Left LOSES
```

### Game Diagram

```
17 ──L(6)──▶ 11 ──R(10)──▶ 1 ──L(?)──▶ ✗
                                       │
                               Right Wins!
```

**Result**: Right wins in 2 moves.

---

## Outcome Classes

In combinatorial game theory, positions are classified into four **outcome classes**:

| Class | Name | Meaning |
|-------|------|---------|
| **N** | Next | The **next player** (first to move) wins with optimal play |
| **P** | Previous | The **previous player** (second to move) wins with optimal play |
| **L** | Left | **Left wins** regardless of who moves first |
| **R** | Right | **Right wins** regardless of who moves first |

### How Classes Are Determined

For each position (L, R, n), we simulate two games:
1. Left goes first → record winner
2. Right goes first → record winner

Then classify:
- Left wins both → **L**
- Right wins both → **R**  
- First player always wins → **N**
- Second player always wins → **P**

---

## Analysis Results

### Overall Distribution (10,000 games)

```
┌─────────────┬───────┬────────────┐
│ Outcome     │ Count │ Percentage │
├─────────────┼───────┼────────────┤
│ N (Next)    │ 1,615    │   16.1%     │
│ P (Previous)│ 1,789 │   17.9%    │
│ L (Left)    │ 3,319 │   33.2%       │
│ R (Right)   │ 3,277 │   32.8%     │
└─────────────┴───────┴────────────┘
```

### Distribution by L vs R Relationship

**When L > R (4,500 games):**
```
N: 741 (16.5%)  │  P: 810 (18.0%)  │  L: 694 (15.4%)  │  R: 2,255 (50.1%)
```
→ Right wins half the games when Left has the larger option!

**When L < R (4,500 games):**
```
N: 738 (16.4%)  │  P: 815 (18.1%)  │  L: 2,333 (51.8%)  │  R: 614 (13.6%)
```
→ Left wins half the games when Left has the smaller option!

**When L = R (1,000 games):**
```
N: 136 (13.6%)  │  P: 164 (16.4%)  │  L: 292 (29.2%)  │  R: 408 (40.8%)
```
→ Slight advantage to Right in symmetric games.

### Distribution by Starting Zone

| Starting Zone | N | P | L | R |
|---------------|---|---|---|---|
| Left's Zone (n mod 5 ∈ {0,1}) | 15.7% | 17.1% | 34.3% | 32.9% |
| Neutral (n mod 5 = 2) | 17.9% | 21.9% | 30.2% | 30.1% |
| Right's Zone (n mod 5 ∈ {3,4}) | 15.7% | 16.7% | 33.6% | 34.0% |

The starting zone has minimal impact on overall win rates.

### Distribution by Counter Size (n = 1 to 10)

| n | N | P | L | R | Notes |
|---|---|---|---|---|-------|
| 1 | 1 | 81 | 9 | 9 | Mostly P (neither can move) |
| 2 | 3 | 64 | 16 | 17 | High P count |
| 3 | 4 | 58 | 23 | 15 | |
| 4 | 7 | 45 | 28 | 20 | |
| 5 | 12 | 33 | 32 | 23 | |
| 6 | 17 | 23 | 35 | 25 | |
| 7 | 19 | 20 | 32 | 29 | |
| 8 | 27 | 13 | 31 | 29 | |
| 9 | 30 | 13 | 26 | 31 | |
| 10 | 40 | 12 | 21 | 27 | Most N positions |

**Pattern**: As n increases, P positions decrease and N positions increase.

---

## Key Findings

### Verified Boundary Conditions

| Condition | Outcome | Games | Explanation |
|-----------|---------|-------|-------------|
| n < min(L, R) | **P** | 285 | Neither player can move first |
| n = L = R | **N** | 10 | Both win immediately if first |
| n = L, L < R | **L** | 45 | Left wins, Right can't move |
| n = R, R < L | **R** | 45 | Right wins, Left can't move |
| L = R = 1, n = 1 | **N** | 1 | First player wins |
| L = R = 1, n > 1 | **R** | 99 | Zone advantage favors Right |

### Special Case: L = R = 1

When both players can only subtract 1:
- n = 1: **N** (first player takes the last counter)
- n > 1: **R** always wins

This happens because the zone multiplier creates a systematic advantage for Right.

---

## Periodicity

### The Discovery

For any fixed L and R, the sequence of outcomes becomes **periodic** after some point.

### Period Table (L, R ∈ {1..5})

```
       R=1   R=2   R=3   R=4   R=5
  L=1    5     5     5     5    35
  L=2    5     6     5     5    35  
  L=3    5     5     6    10    40
  L=4    5    10    25    40    50
  L=5   35    35    40    50    30
```

### Why Periodicity Occurs

Three cycles interact:

1. **Zone Cycle** (period 5): LZ → LZ → Neutral → RZ → RZ → repeat
2. **Left's Move Cycle** (related to L): Where Left lands mod L
3. **Right's Move Cycle** (related to R): Where Right lands mod R

The **doubling rule** creates cascading effects that make the actual period more complex than lcm(L, R, 5).

### Example: L = 3, R = 5 (Period = 40)

```
n=1-10:   P P L L N L N P L R
n=11-20:  R L R P L P P N N N  
n=21-30:  N N P P P R P N R L
n=31-40:  L R L N R N N P P P  ← End of period
n=41-50:  P P N N N L N P L R  ← Pattern repeats
```

### Practical Use

Once you know the period, predict any outcome:
```
outcome(n) = outcome((n - 1) mod Period + 1)

Example: L=3, R=5, Period=40
outcome(147) = outcome((147-1) mod 40 + 1) = outcome(27) = P
```

---

## The Smaller Option Advantage

### Key Insight

**The player with the smaller subtraction option tends to win more often!**

| Condition | Dominant Winner | Win Rate |
|-----------|-----------------|----------|
| L < R | Left | 51.8% |
| L > R | Right | 50.1% |

### Why This Happens

The zone multiplier (×2) benefits players who can make **more granular moves**:

1. **More zone interactions**: Smaller options cross zone boundaries more frequently
2. **Better control**: Fine-grained moves allow better positioning
3. **Multiplier exploitation**: More opportunities to trigger (or avoid) the doubling rule

---


### Running the Analysis

```bash
cd src/
go run main.go
```

This will:
1. Generate all 10,000 game combinations
2. Simulate each game twice (Left first, Right first)
3. Classify outcomes and output to CSV

### Output Format

The CSV file contains:
```
L Option, R Option, Counter, Outcome Class, LGF Winner, RGF Winner
1, 1, 1, N, L, R
1, 1, 2, R, R, R
...
```

Where:
- **LGF Winner**: Winner when Left Goes First
- **RGF Winner**: Winner when Right Goes First

## License
MIT License - See LICENSE file for details.