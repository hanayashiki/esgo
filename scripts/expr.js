const pred = {
  '+': 1,
  '*': 2,
}

const code = '1*1*2';

const Lexer = () => {
  let pos = 0;
  let currentToken = undefined;

  const next = () => {
    while (code[pos] === ' ') {
      pos++;
    }

    currentToken = code[pos];
    pos++;

    return currentToken;
  };

  const peek = () => currentToken;

  return {
    next,
    peek,
  };
}

const lexer = Lexer();
lexer.next();

const parseExpr = (level = 0) => {
  const prefix = parsePrefix();

  return parseSuffix(prefix, level);
};

const parsePrefix = () => {
  if (lexer.peek() >= '0' && lexer.peek() <= '9') {
    const value = parseInt(lexer.peek())
    lexer.next();
    return {
      type: 'NumericLiteral',
      value,
    }
  }
  throw new Error();
}

const parseSuffix = (left, level) => {
  while (true) {
    const op = lexer.peek();
    
    if (op === undefined || level >= pred[op]) {
      return left;
    }
    if (pred[op] === undefined) {
      throw new Error('unexpected operator ' + pred[op])
    }
    lexer.next();

    left = {
      type: 'BinaryExpression',
      left,
      operator: op,
      right: parseExpr(pred[op]),
    }
  }
};

console.log(JSON.stringify(parseExpr(), null, 2));