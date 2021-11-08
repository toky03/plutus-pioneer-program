import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { LoveLaceOffer, Wallet } from '../model/model';

const API_ROOT = environment.apiBaseUrl;
@Injectable({
  providedIn: 'root'
})
export class IntegrationService {

  constructor(private httpClient: HttpClient) { }


  public readWallets(): Observable<Wallet[]> {
    return this.httpClient.get<Wallet[]>(`${API_ROOT}/wallets`)
  }
  public readFunds(id: string): Observable<Wallet> {
    return this.httpClient.get<Wallet>(`${API_ROOT}/${id}/funds`);
  }

  public offerLovelace(id: string, lovelaces: LoveLaceOffer): Observable<void> {
    return this.httpClient.post<void>(`${API_ROOT}/${id}/offer`, lovelaces);
  }

  public retreive(id: string): Observable<void> {
    return this.httpClient.post<void>(`${API_ROOT}/${id}/retrieve`, null);
  }

  public use(id: string): Observable<void> {
    return this.httpClient.post<void>(`${API_ROOT}/${id}/use`, null)
  }
}
