package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Depositar(valor float64) (bool, string) {
	if valor <= 0 {
		return false, "Não é possível depositar valores menores ou igual a zero"
	}
	c.saldo += valor
	return true, "Deposito efetuado com sucesso"
}

func (c *ContaCorrente) Sacar(valor float64) (bool, string) {
	if valor > 0 && c.saldo >= valor {
		c.saldo -= valor
		return true, "Saque efetuado com sucesso"
	}
	return false, "Não foi possível efetuar o saque"
}

func (c *ContaCorrente) Transferir(contaDestino *ContaCorrente, valor float64) (bool, string) {
	if valor <= 0 {
		return false, "Não é possível transferir valores menores ou igual a zero"
	}
	status, mensagem := c.Sacar(valor)
	if status {
		status, mensagem = contaDestino.Depositar(valor)
		if status {
			return true, "Transferencia realizada com sucesso"
		}
	}

	return false, mensagem
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
