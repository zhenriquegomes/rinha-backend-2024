CREATE TABLE public.clientes (
    id INTEGER NOT NULL PRIMARY KEY,
    limite INTEGER NOT NULL,
    saldo_inicial INTEGER NOT NULL,
    saldo_atual INTEGER NOT NULL
);

INSERT INTO public.clientes 
    (id, limite, saldo_inicial, saldo_atual)
VALUES
    (1, 100000, 0, 0),
    (2, 80000, 0, 0),
    (3, 1000000, 0, 0),
    (4, 10000000, 0, 0),
    (5, 500000, 0, 0);

CREATE TABLE public.transacoes (
    id SERIAL NOT NULL PRIMARY KEY,
    client_id INTEGER NOT NULL REFERENCES public.clientes(id),
    valor INTEGER NOT NULL,
    tipo VARCHAR(1) NOT NULL,
    descricao VARCHAR(10) NOT NULL,
    dt_referencia TIMESTAMP NOT NULL
);