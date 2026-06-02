export function computeShift(shift, angle){
    const angleRad = angle * Math.PI / 180;
    const leftShift = shift * Math.cos(angleRad);
    const topShift = shift * Math.sin(angleRad);
    return [leftShift, topShift]
}

export function shiftLine(left1, top1, left2, top2, shift, angle){
    const [leftShift, topShift] = computeShift(shift, angle)
    return [left1+leftShift, top1+topShift, left2-leftShift, top2-topShift]
}

export function computeLine(left1, top1, left2, top2){
    let left = left2 - left1
    let top = top2 - top1
    let hypotenuse = Math.sqrt(Math.pow(left, 2) + Math.pow(top, 2))

    let radians = Math.acos(left/hypotenuse);
    let degrees = radians * (180 / Math.PI);
    if (top < 0) {
        degrees = -degrees
    }
    return [left, top, hypotenuse, degrees]
}