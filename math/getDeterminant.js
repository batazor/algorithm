function isRightArrayNumber(A) { return Number.isInteger(Math.sqrt(A)) }
function isCoolMatrix(A) { return A.length === 4; }
function getDeterminant(A) { return (A[0] * A[3]) - (A[1] * A[2]); }

function getSubMatrix(A, step, iA, jA) {
  const subMatrix = [];
  
  A.forEach((item, index) => {
    const isI = (index % step) === 0;
    const isJ = parseInt(index / step) === jA;
    
    !isI && !isJ && subMatrix.push(item);
  });
  
  return subMatrix;
}

function getMatrix(A) {
  if (isCoolMatrix(A)) { return getDeterminant(A); }

  const step = Math.sqrt(A.length);
  let SUM = 0;
  
  for(let counter = 0; counter < step; counter++) {
    const index = counter * step;
    const i = 1;
    const j = counter + 1;
    const M = getSubMatrix(A, step, i, counter);
    
    SUM += A[index] * Math.pow(-1, i + j) * getMatrix(M);
  }

  return SUM;
}

export default getMatrix;
