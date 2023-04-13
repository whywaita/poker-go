type f = () => { ok: true, handCards: string, notUsedCards: string, hand: string } | { ok: false, message: string };

declare var poker: {
    generateHands: f,
}