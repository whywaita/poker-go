import Card, {CardType, Suit, Rank} from '../components/Card'
import Script from 'next/script'
import {useCallback, useState} from "react";

export default function Home() {
    const [cards, setCards] = useState<CardType[]>([]);
    const [notUsedCards, setNotUsedCards] = useState<CardType[]>([]);
    const [hand, setHand] = useState<string>("Loading...");

    const generate = useCallback(() => {
        let result = poker.generateHands();
        if (!result) {console.error("No result"); return;}
        if (result.ok) {
            let newData: CardType[] = [];
            let newHand: string = "";
            try {
                newData = JSON.parse(result.handCards);
                newHand = result.hand;
            }  catch (e) {
                console.error("Invalid JSON (cards)", e);
                return;
            }
            if (newData.length !== 5) {
                console.error("Invalid number of cards");
                return;
            }

            let newUnUsedData: CardType[];
            try {
                newUnUsedData = JSON.parse(result.notUsedCards);
            } catch (e) {
                console.error("Invalid JSON (Unused cards)", e);
                return;
            }

            setCards({...newData});
            setNotUsedCards({...newUnUsedData});
            setHand(newHand);
        } else {
            console.error(result.message);
        }
    },[])

    const CardComponent = () => {
        if (!cards) { return <div>Loading...</div> }

        if (!(cards[0] as CardType) || !(cards[1] as CardType) || !(cards[2] as CardType) || !(cards[3] as CardType) || !(cards[4] as CardType)) {
            return <div>Loading...</div>
        }

        const c1: CardType = {
            suit: cards[0].suit as Suit,
            rank: cards[0].rank as Rank
        }
        const c2: CardType = {
            suit: cards[1].suit as Suit,
            rank: cards[1].rank as Rank
        }
        const c3: CardType = {
            suit: cards[2].suit as Suit,
            rank: cards[2].rank as Rank
        }
        const c4: CardType = {
            suit: cards[3].suit as Suit,
            rank: cards[3].rank as Rank
        }
        const c5: CardType = {
            suit: cards[4].suit as Suit,
            rank: cards[4].rank as Rank
        }

        return (
            <div className={"flex"}>
                <Card suit={c1.suit} rank={c1.rank} />
                <Card suit={c2.suit} rank={c2.rank} />
                <Card suit={c3.suit} rank={c3.rank} />
                <Card suit={c4.suit} rank={c4.rank} />
                <Card suit={c5.suit} rank={c5.rank} />
            </div>
        )
    }

    const NotUsedCardComponent = () => {
        if (!notUsedCards) { return <div>Loading...</div> }

        if (!(notUsedCards[0] as CardType) || !(notUsedCards[1] as CardType)) {
            return <div>Loading...</div>
        }

        const nc1: CardType = {
            suit: notUsedCards[0].suit as Suit,
            rank: notUsedCards[0].rank as Rank
        }
        const nc2: CardType = {
            suit: notUsedCards[1].suit as Suit,
            rank: notUsedCards[1].rank as Rank
        }

        return (
            <div className={"flex"}>
                <Card suit={nc1.suit} rank={nc1.rank} />
                <Card suit={nc2.suit} rank={nc2.rank} />
            </div>
        )
    }

    return (
      <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <Script id="exec-wasm" src="/poker-go/wasm_exec.js" onLoad={() => {
            // @ts-ignore
            const go = new Go();
            WebAssembly.instantiateStreaming(fetch("/poker-go/pokergo.wasm"), go.importObject).then((result) => {
                go.run(result.instance);
            });
        }}/>

        <div className="navbar bg-base-100">
          <a className="btn btn-ghost normal-case text-xl">Poker Annotation</a>
        </div>

        <CardComponent />
        <NotUsedCardComponent />

        <p className={"text-3xl"}>Is</p>

        <p id={"hand"} className={"text-5xl"}>{hand}</p>

        <button className={"btn normal-case h-10 text-2xl"} onClick={generate}>Next</button>

        {/*<div className={"flex flex-wrap gap-3"}>*/}
        {/*    <button className="btn btn-info h-20 text-4xl">OK</button>*/}
        {/*    <button className="btn btn-error h-20 text-4xl">NG</button>*/}
        {/*</div>*/}
    </main>
  )
}
