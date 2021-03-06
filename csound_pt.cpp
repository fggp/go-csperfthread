#include <csound/csound.hpp>
#include <csound/csPerfThread.hpp>
#include "csound_pt.h"

extern "C" {

//#include "_cgo_export.h"
  
Cpt NewCsoundPT(CSOUND *csound)
{
  CsoundPerformanceThread *pt = new CsoundPerformanceThread(csound);
  return (void *)pt;
}

void DeleteCsoundPT(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  delete cpt;
}

int CsoundPTisRunning(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  return cpt->isRunning();
}

/*////////////////////////////////////////////////////////////*/

extern void goPTprocessCB(void *);

void CsoundPTProcessCB(void *cbData)
{
  goPTprocessCB(cbData);
}

void CsoundPTsetProcessCB(Cpt pt, void *cbData)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->SetProcessCallback(CsoundPTProcessCB, cbData);
}

/*////////////////////////////////////////////////////////////*/

CSOUND *CsoundPTgetCsound(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  return cpt->GetCsound();
}

int CsoundPTgetStatus(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  return cpt->GetStatus();
}

void CsoundPTplay(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->Play();
}

void CsoundPTpause(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->Pause();
}

void CsoundPTtogglePause(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->TogglePause();
}

void CsoundPTstop(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->Stop();
}

void CsoundPTrecord(Cpt pt, const char *filename, int samplebits, int numbufs)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  std::string fname(filename);
  cpt->Record(fname, samplebits, numbufs);
}

void CsoundPTstopRecord(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->StopRecord();
}

void CsoundPTscoreEvent(Cpt pt, int absp2mode, char opcod, int pcnt, MYFLT *p)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->ScoreEvent(absp2mode, opcod, pcnt, p);
}

void CsoundPTinputMessage(Cpt pt, const char *s)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->InputMessage(s);
}

void CsoundPTsetScoreOffsetSeconds(Cpt pt, double timeVal)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->SetScoreOffsetSeconds(timeVal);
}

int CsoundPTjoin(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  return cpt->Join();
}

void CsoundPTflushMessageQueue(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->FlushMessageQueue();
}

} // extern "C"
