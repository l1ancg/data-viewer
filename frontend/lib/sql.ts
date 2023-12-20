import { Column, Parser, Select } from 'node-sql-parser';
import { ParseResult } from '@/types/sql';

const parser = new Parser();

/**
 * only support select
 * @param inputSql select sql
 * @param type MySQL|...
 */
const parse = (inputSql: string, type: string): ParseResult => {
  try {
    let ast = parser.astify(inputSql, { database: type });
    if (Array.isArray(ast)) {
      ast = ast[0];
    }
    let selectAst = ast as Select;
    if (selectAst.columns.length > 0) {
      fillColumnAs(selectAst.columns);
    }
    // let columnList: string[] = [];
    // if (selectAst.columns !== '*') {
    //   columnList = selectAst.columns.map((c: Column) => c.as);
    // if (Array.isArray(e) || e === "*") {
    //   console.log(e);
    // } else {
    //   throw new Error("parse error");
    // }
    return {
      sql: parser.sqlify(ast),
      whereList: selectAst.where ? parseWhere(selectAst.where) : [],
      columnList:
        typeof selectAst.columns === 'string' && selectAst.columns === '*'
          ? []
          : selectAst.columns.map((c) => c.as),
    };
  } catch (e: any) {
    console.log(e);

    if ('location' in e) {
      let ee = e.location;
      let start = { ...ee.start };
      let end = { ...ee.end };
      return {
        error: { start, end },
      };
    }
  }
  throw new Error('parse error');
};

const fillColumnAs = (columns: any[] | Column[] | '*'): void => {
  for (let i = 0; i < columns.length; i++) {
    if (columns[i].expr.type === 'column_ref') {
      columns[i].as = columns[i].expr.column;
    }
  }
};

const parseWhere = (where: any): string[] => {
  const whereList: string[] = [];
  switch (where.operator) {
    case '=':
      if (where.right.type === 'origin') {
        whereList.push(
          (where.left.table ? where.left.table + '.' : '') + where.left.column
        );
      }
      break;
    case 'AND':
      const lwl = parseWhere(where.left);
      const rwl = parseWhere(where.right);
      whereList.push(...lwl, ...rwl);
      break;
    case 'IN':
      whereList.push(
        (where.left.table ? where.left.table + '.' : '') +
          where.left.column +
          ' in ' +
          where.right.value[0].ast.from[0].table
      );
      break;
  }
  return whereList;
};

export default parse;
